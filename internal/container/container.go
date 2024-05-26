package container

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	"github.com/dcwk/gophermart/internal/config"
	"github.com/dcwk/gophermart/internal/repositories"
	"github.com/dcwk/gophermart/internal/services"
)

type Container struct {
	conf    *config.ServerConf
	DB_     *pgxpool.Pool
	Logger_ *zap.Logger

	UserRepository_  repositories.UserRepository
	OrderRepository_ repositories.OrderRepository

	RegisterUserService_ *services.RegisterUserService
	AuthUserService_     *services.AuthUserService
}

func NewContainer(conf *config.ServerConf) *Container {
	return &Container{
		conf: conf,
	}
}

func (c *Container) DB() *pgxpool.Pool {
	if c.DB_ == nil {
		conf, err := pgxpool.ParseConfig(c.conf.DatabaseDSN)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to parse database DSN: %v\n", err)
			os.Exit(1)
		}
		if c.conf.DatabaseMaxConnections > 0 {
			conf.MaxConns = int32(c.conf.DatabaseMaxConnections)
		}

		c.DB_, err = pgxpool.NewWithConfig(context.Background(), conf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
			os.Exit(1)
		}
	}

	return c.DB_
}

func (c *Container) Logger() *zap.Logger {
	if c.Logger_ == nil {
		lvl, err := zap.ParseAtomicLevel(c.conf.LogLevel)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to parse log level: %v\n", err)
			os.Exit(1)
		}

		cfg := zap.NewProductionConfig()
		cfg.Level = lvl
		zl, err := cfg.Build()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to initialize logger: %v\n", err)
			os.Exit(1)
		}

		c.Logger_ = zl
	}

	return c.Logger_
}

func (c *Container) UserRepository() repositories.UserRepository {
	if c.UserRepository_ == nil {
		c.UserRepository_ = repositories.NewUserRepository(c.DB())
	}

	return c.UserRepository_
}

func (c *Container) OrderRepository() repositories.OrderRepository {
	if c.OrderRepository_ == nil {
		c.OrderRepository_ = repositories.NewOrderRepository(c.DB())
	}

	return c.OrderRepository_
}

func (c *Container) RegisterUserService() *services.RegisterUserService {
	if c.RegisterUserService_ == nil {
		c.RegisterUserService_ = services.NewRegisterUserService(c.UserRepository())
	}

	return c.RegisterUserService_
}

func (c *Container) AuthUserService() *services.AuthUserService {
	if c.AuthUserService_ == nil {
		c.AuthUserService_ = services.NewAuthService(c.UserRepository())
	}

	return c.AuthUserService_
}
