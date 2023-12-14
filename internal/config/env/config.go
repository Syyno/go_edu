package envconfig

import (
	"fmt"
	"os"
	"users/internal/common"

	"github.com/joho/godotenv"
)

type Provider interface {
	Get(key string) (string, error)
}

// Configuration ...
type Configuration struct {
	provider Provider
}

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return common.NewErrorf(common.ErrorCodeUnknown, "loading env var file")
	}

	return nil
}

// New ...
func New(provider Provider) *Configuration {
	return &Configuration{
		provider: provider,
	}
}

// Get returns the value from environment variable `<key>`. When an environment variable `<key>_SECURE` exists
// the provider is used for getting the value.
func (c *Configuration) Get(key string) (string, error) {
	res := os.Getenv(key)
	valSecret := os.Getenv(fmt.Sprintf("%s_SECURE", key))

	if valSecret != "" {
		valSecretRes, err := c.provider.Get(valSecret)
		if err != nil {
			return "", common.WrapErrorf(err, common.ErrorCodeInvalidArgument, "provider.Get")
		}

		res = valSecretRes
	}

	return res, nil
}
