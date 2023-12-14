package envconfig

import (
	"net"

	"github.com/pkg/errors"
)

const (
	grpcHostEnvName = "GRPC_HOST"
	grpcPortEnvName = "GRPC_PORT"
)

type GRPCConfig interface {
	Address() string
}

type grpcConfig struct {
	host string
	port string
}

func NewGRPCConfig(conf *Configuration) (GRPCConfig, error) {
	//host := os.Getenv(grpcHostEnvName)
	hostValue, err := conf.Get(grpcHostEnvName)
	if err != nil {
		return nil, errors.New("grpc host not found")
	}

	portValue, err := conf.Get(grpcPortEnvName)
	if err != nil {
		return nil, errors.New("grpc port not found")
	}

	return &grpcConfig{
		host: hostValue,
		port: portValue,
	}, nil
}

func (cfg *grpcConfig) Address() string {
	//return ":50051"
	return net.JoinHostPort(cfg.host, cfg.port)
}
