package envconfig

import (
	"errors"
	"net"
)

const (
	httpHostEnvName = "HTTP_HOST"
	httpPortEnvName = "HTTP_PORT"
)

type HTTPConfig interface {
	Address() string
}

type httpConfig struct {
	host string
	port string
}

func NewHTTPConfig(conf *Configuration) (HTTPConfig, error) {
	hostValue, err := conf.Get(httpHostEnvName)
	if err != nil {
		return nil, errors.New("http host not found")
	}

	portValue, err := conf.Get(httpPortEnvName)
	if err != nil {
		return nil, errors.New("http port not found")
	}

	return &httpConfig{
		host: hostValue,
		port: portValue,
	}, nil
}

func (cfg *httpConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
