package container

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	"github.com/dcwk/gophermart/internal/config"
	"github.com/dcwk/gophermart/internal/repositories"
	"github.com/dcwk/gophermart/internal/usecase"
)

type Container struct {
	conf    *config.ServerConf
	DB_     *pgxpool.Pool
	Logger_ *zap.Logger

	UserRepository_        repositories.UserRepository
	OrderRepository_       repositories.OrderRepository
	UserBalanceRepository_ repositories.UserBalanceRepository
	AccrualRepository_     repositories.AccrualRepository
	WithdrawalRepository_  repositories.WithdrawalRepository

	RegisterUserHandler_    *usecase.RegisterUserHandler
	AuthUserHandler_        *usecase.AuthUserHandler
	GetOrdersHandler_       *usecase.GetOrdersHandler
	GetUserBalanceHandler_  *usecase.GetUserBalanceHandler
	GetWithdrawalsHandler_  *usecase.GetWithdrawalsHandler
	CreateOrderHandler_     *usecase.CreateOrderHandler
	LoadOrderHandler_       *usecase.LoadOrderHandler
	WithdrawRequestHandler_ *usecase.WithdrawRequestHandler
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

func (c *Container) UserBalanceRepository() repositories.UserBalanceRepository {
	if c.UserBalanceRepository_ == nil {
		c.UserBalanceRepository_ = repositories.NewUserBalanceRepository(c.DB())
	}

	return c.UserBalanceRepository_
}

func (c *Container) AccrualRepository() repositories.AccrualRepository {
	if c.AccrualRepository_ == nil {
		c.AccrualRepository_ = repositories.NewAccrualRepository(c.DB())
	}

	return c.AccrualRepository_
}

func (c *Container) WithdrawalRepository() repositories.WithdrawalRepository {
	if c.WithdrawalRepository_ == nil {
		c.WithdrawalRepository_ = repositories.NewWithdrawalRepository(c.DB())
	}

	return c.WithdrawalRepository_
}

func (c *Container) RegisterUserHandler() *usecase.RegisterUserHandler {
	if c.RegisterUserHandler_ == nil {
		c.RegisterUserHandler_ = usecase.NewRegisterUserHandler(c.UserRepository(), c.UserBalanceRepository())
	}

	return c.RegisterUserHandler_
}

func (c *Container) AuthUserHandler() *usecase.AuthUserHandler {
	if c.AuthUserHandler_ == nil {
		c.AuthUserHandler_ = usecase.NewAuthHandler(c.UserRepository())
	}

	return c.AuthUserHandler_
}

func (c *Container) GetOrdersHandler() *usecase.GetOrdersHandler {
	if c.GetOrdersHandler_ == nil {
		c.GetOrdersHandler_ = usecase.NewGetOrdersHandler(c.UserRepository(), c.OrderRepository())
	}

	return c.GetOrdersHandler_
}

func (c *Container) GetUserBalanceHandler() *usecase.GetUserBalanceHandler {
	if c.GetUserBalanceHandler_ == nil {
		c.GetUserBalanceHandler_ = usecase.NewGetUserBalanceHandler(c.UserRepository(), c.UserBalanceRepository())
	}

	return c.GetUserBalanceHandler_
}

func (c *Container) GetWithdrawalsHandler() *usecase.GetWithdrawalsHandler {
	if c.GetWithdrawalsHandler_ == nil {
		c.GetWithdrawalsHandler_ = usecase.NewGetWithdrawalsHandler(c.UserRepository(), c.WithdrawalRepository())
	}

	return c.GetWithdrawalsHandler_
}

func (c *Container) CreateOrderHandler() *usecase.CreateOrderHandler {
	if c.CreateOrderHandler_ == nil {
		c.CreateOrderHandler_ = usecase.NewCreateOrderHandler(
			c.Logger(),
			c.UserRepository(),
			c.OrderRepository(),
			c.AccrualRepository(),
			c.UserBalanceRepository(),
		)
	}

	return c.CreateOrderHandler_
}

func (c *Container) LoadOrderHandler() *usecase.LoadOrderHandler {
	if c.LoadOrderHandler_ == nil {
		c.LoadOrderHandler_ = usecase.NewLoadOrderHandler(
			c.conf.AccrualSystemAddress,
			c.Logger(),
		)
	}

	return c.LoadOrderHandler_
}

func (c *Container) WithdrawRequestHandler() *usecase.WithdrawRequestHandler {
	if c.WithdrawRequestHandler_ == nil {
		c.WithdrawRequestHandler_ = usecase.NewWithdrawRequestHandler(
			c.Logger(),
			c.UserRepository(),
			c.UserBalanceRepository(),
			c.OrderRepository(),
			c.WithdrawalRepository(),
		)
	}

	return c.WithdrawRequestHandler_
}
