package container

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	"github.com/dcwk/gophermart/internal/config"
	"github.com/dcwk/gophermart/internal/repositories"
	"github.com/dcwk/gophermart/internal/use_case"
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

	RegisterUserService_    *use_case.RegisterUserService
	AuthUserService_        *use_case.AuthUserService
	GetOrdersService_       *use_case.GetOrdersService
	GetUserBalanceService_  *use_case.GetUserBalanceService
	GetWithdrawalsService_  *use_case.GetWithdrawalsService
	CreateOrderService_     *use_case.CreateOrderService
	LoadOrderService_       *use_case.LoadOrderService
	WithdrawRequestService_ *use_case.WithdrawRequestService
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

func (c *Container) RegisterUserService() *use_case.RegisterUserService {
	if c.RegisterUserService_ == nil {
		c.RegisterUserService_ = use_case.NewRegisterUserService(c.UserRepository(), c.UserBalanceRepository())
	}

	return c.RegisterUserService_
}

func (c *Container) AuthUserService() *use_case.AuthUserService {
	if c.AuthUserService_ == nil {
		c.AuthUserService_ = use_case.NewAuthService(c.UserRepository())
	}

	return c.AuthUserService_
}

func (c *Container) GetOrdersService() *use_case.GetOrdersService {
	if c.GetOrdersService_ == nil {
		c.GetOrdersService_ = use_case.NewGetOrdersService(c.UserRepository(), c.OrderRepository())
	}

	return c.GetOrdersService_
}

func (c *Container) GetUserBalanceService() *use_case.GetUserBalanceService {
	if c.GetUserBalanceService_ == nil {
		c.GetUserBalanceService_ = use_case.NewGetUserBalanceService(c.UserRepository(), c.UserBalanceRepository())
	}

	return c.GetUserBalanceService_
}

func (c *Container) GetWithdrawalsService() *use_case.GetWithdrawalsService {
	if c.GetWithdrawalsService_ == nil {
		c.GetWithdrawalsService_ = use_case.NewGetWithdrawalsService(c.UserRepository(), c.WithdrawalRepository())
	}

	return c.GetWithdrawalsService_
}

func (c *Container) CreateOrderService() *use_case.CreateOrderService {
	if c.CreateOrderService_ == nil {
		c.CreateOrderService_ = use_case.NewCreateOrderService(
			c.Logger(),
			c.UserRepository(),
			c.OrderRepository(),
			c.AccrualRepository(),
			c.UserBalanceRepository(),
		)
	}

	return c.CreateOrderService_
}

func (c *Container) LoadOrderService() *use_case.LoadOrderService {
	if c.LoadOrderService_ == nil {
		c.LoadOrderService_ = use_case.NewLoadOrderService(
			c.conf.AccrualSystemAddress,
			c.Logger(),
		)
	}

	return c.LoadOrderService_
}

func (c *Container) WithdrawRequestService() *use_case.WithdrawRequestService {
	if c.WithdrawRequestService_ == nil {
		c.WithdrawRequestService_ = use_case.NewWithdrawRequestService(
			c.Logger(),
			c.UserRepository(),
			c.UserBalanceRepository(),
			c.OrderRepository(),
			c.WithdrawalRepository(),
		)
	}

	return c.WithdrawRequestService_
}
