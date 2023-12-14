package envconfig

import (
	"net"

	"github.com/pkg/errors"
)

const (
	swaggerHostEnvName = "SWAGGER_HOST"
	swaggerPortEnvName = "SWAGGER_PORT"
)

type SwaggerConfig interface {
	Address() string
}

type swaggerConfig struct {
	host string
	port string
}

func NewSwaggerConfig(conf *Configuration) (SwaggerConfig, error) {
	hostValue, err := conf.Get(swaggerHostEnvName)
	if err != nil {
		return nil, errors.New("swagger host not found")
	}

	portValue, err := conf.Get(swaggerPortEnvName)
	if err != nil {
		return nil, errors.New("swagger port not found")
	}

	return &swaggerConfig{
		host: hostValue,
		port: portValue,
	}, nil
}

func (cfg *swaggerConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
