package vault

import (
	"fmt"
	"os"
	"strings"
	"users/internal/common"

	"github.com/hashicorp/vault/api"
)

const (
	vaultPathEnvName    = "VAULT_PATH"
	vaultTokenEnvName   = "VAULT_TOKEN"
	vaultAddressEnvName = "VAULT_ADDRESS"
)

type Provider struct {
	path    string
	client  *api.Logical
	results map[string]map[string]string
}

func New(token, addr, path string) (*Provider, error) {
	config := &api.Config{
		Address: addr,
	}

	client, err := api.NewClient(config)
	if err != nil {
		return nil, common.WrapErrorf(err, common.ErrorCodeUnknown, "api.NewClient")
	}

	client.SetToken(token)

	return &Provider{
		path:    path,
		client:  client.Logical(),
		results: make(map[string]map[string]string),
	}, nil
}

func (p *Provider) Get(v string) (string, error) {
	split := strings.Split(v, ":")
	if len(split) == 1 {
		return "", common.NewErrorf(common.ErrorCodeUnknown, "missing key value")
	}

	pathSecret := split[0]
	key := split[1]

	res, ok := p.results[pathSecret]
	if ok {
		val, ok := res[key]
		if !ok {
			return "", common.NewErrorf(common.ErrorCodeUnknown, "key not found in cached data")
		}

		return val, nil
	}

	secret, err := p.client.Read(fmt.Sprintf("%s/data/%s", p.path, pathSecret))
	if err != nil {
		return "", common.WrapErrorf(err, common.ErrorCodeUnknown, "reading")
	}

	if secret == nil {
		return "", common.NewErrorf(common.ErrorCodeUnknown, "secret not found")
	}

	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		return "", common.NewErrorf(common.ErrorCodeUnknown, "invalid data in secret")
	}

	secrets := make(map[string]string)

	for k, v := range data {
		val, ok := v.(string)
		if !ok {
			return "", common.NewErrorf(common.ErrorCodeUnknown, "secret value in data is not string")
		}

		secrets[k] = val
	}

	val, ok := secrets[key]
	if !ok {
		return "", common.NewErrorf(common.ErrorCodeUnknown, "key not found in retrieved data")
	}

	p.results[pathSecret] = secrets

	return val, nil
}

func NewVaultProvider() (*Provider, error) {
	vaultPath := os.Getenv(vaultPathEnvName)
	vaultToken := os.Getenv(vaultTokenEnvName)
	vaultAddress := os.Getenv(vaultAddressEnvName)

	provider, err := New(vaultToken, vaultAddress, vaultPath)
	if err != nil {
		return nil, common.WrapErrorf(err, common.ErrorCodeUnknown, "vault.New ")
	}

	return provider, nil
}
