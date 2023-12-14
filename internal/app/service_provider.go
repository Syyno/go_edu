package app

import (
	"context"
	"log"
	"users/internal/api/v1/access"
	"users/internal/api/v1/auth"
	"users/internal/api/v1/user"
	"users/internal/app/closer"
	envconfig "users/internal/config/env"
	"users/internal/repository"
	"users/internal/repository/postgres/db"
	"users/internal/repository/postgres/db/transaction"
	pgRepo "users/internal/repository/postgres/user"
	hstRepo "users/internal/repository/postgres/user_update_history"
	accessService "users/internal/service/access"
	accessServiceImplementation "users/internal/service/access/handlers"
	authService "users/internal/service/auth"
	authServiceImplementation "users/internal/service/auth/handlers"
	userService "users/internal/service/user"
	userServiceImplementation "users/internal/service/user/handlers"
)

type serviceProvider struct {
	grpcConfig    envconfig.GRPCConfig
	pgConfig      envconfig.PGConfig
	httpConfig    envconfig.HTTPConfig
	swaggerConfig envconfig.SwaggerConfig
	authConfig    envconfig.AuthConfig

	userRepository        repository.UserRepository
	userUpdateHistoryRepo repository.UserHistoryRepository

	userService   userService.UserService
	authService   authService.AuthService
	accessService accessService.AccessService

	userImpl   *user.Implementation
	authImpl   *auth.Implementation
	accessImpl *access.Implementation

	dbClient  db.Client
	txManager db.TxManager
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GRPCConfig(conf *envconfig.Configuration) envconfig.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := envconfig.NewGRPCConfig(config)
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) HTTPConfig(conf *envconfig.Configuration) envconfig.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := envconfig.NewHTTPConfig(config)
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) PGConfig(conf *envconfig.Configuration) envconfig.PGConfig {
	if s.pgConfig == nil {
		cfg, err := envconfig.NewPGConfig(config)
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) AuthConfig(conf *envconfig.Configuration) envconfig.AuthConfig {
	if s.authConfig == nil {
		cfg, err := envconfig.NewAuthConfig(config)
		if err != nil {
			log.Fatalf("failed to get auth config: %s", err.Error())
		}

		s.authConfig = cfg
	}

	return s.authConfig
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = pgRepo.NewUserRepository(s.DBClient(ctx))
	}

	return s.userRepository
}

func (s *serviceProvider) UserUpdateHistoryRepository(ctx context.Context) repository.UserHistoryRepository {
	if s.userUpdateHistoryRepo == nil {
		s.userUpdateHistoryRepo = hstRepo.NewUserUpdateHistoryRepo(s.DBClient(ctx))
	}

	return s.userUpdateHistoryRepo
}

func (s *serviceProvider) UserService(ctx context.Context) userService.UserService {
	if s.userService == nil {
		s.userService = userServiceImplementation.NewUserService(
			s.UserRepository(ctx),
			s.UserUpdateHistoryRepository(ctx),
			s.TxManager(ctx),
		)
	}
	return s.userService
}

func (s *serviceProvider) AuthService(ctx context.Context) authService.AuthService {
	if s.authService == nil {
		s.authService = authServiceImplementation.NewAuthService(
			s.AuthConfig(config),
			s.UserService(ctx))
	}
	return s.authService
}

func (s *serviceProvider) AccessService(_ context.Context) accessService.AccessService {
	if s.accessService == nil {
		s.accessService = accessServiceImplementation.NewAccessService(s.AuthConfig(config))
	}
	return s.accessService
}

func (s *serviceProvider) AccessImpl(ctx context.Context) *access.Implementation {
	if s.accessImpl == nil {
		s.accessImpl = access.NewImplementation(s.AccessService(ctx))
	}

	return s.accessImpl
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := db.New(ctx, s.PGConfig(config).DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v. DSN = %v", err, s.PGConfig(config).DSN())
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) UserImpl(ctx context.Context) *user.Implementation {
	if s.userImpl == nil {
		s.userImpl = user.NewImplementation(s.UserService(ctx))
	}

	return s.userImpl
}

func (s *serviceProvider) AuthImpl(ctx context.Context) *auth.Implementation {
	if s.authImpl == nil {
		s.authImpl = auth.NewImplementation(s.AuthService(ctx))
	}

	return s.authImpl
}

func (s *serviceProvider) SwaggerConfig() envconfig.SwaggerConfig {
	if s.swaggerConfig == nil {
		cfg, err := envconfig.NewSwaggerConfig(config)
		if err != nil {
			log.Fatalf("failed to get swagger config: %s", err.Error())
		}

		s.swaggerConfig = cfg
	}

	return s.swaggerConfig
}
