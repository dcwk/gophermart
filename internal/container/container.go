package container

import (
	"database/sql"

	"github.com/dcwk/gophermart/internal/config"
	"github.com/dcwk/gophermart/internal/repositories"
	"github.com/dcwk/gophermart/internal/services"
)

type Container struct {
	conf *config.ServerConf
	DB_  *sql.DB

	UserRepository_ repositories.UserRepository

	RegisterUserService_ *services.RegisterUserService
	AuthUserService_     *services.AuthUserService
}

func NewContainer(conf *config.ServerConf) *Container {
	return &Container{
		conf: conf,
	}
}

func (c *Container) DB() *sql.DB {
	if c.DB_ == nil {
		var err error
		c.DB_, err = sql.Open("pgx", c.conf.DatabaseDSN)
		if err != nil {
			panic(err)
		}
	}

	return c.DB_
}

func (c *Container) UserRepository() repositories.UserRepository {
	if c.UserRepository_ == nil {
		c.UserRepository_ = repositories.NewUserRepository(c.DB())
	}

	return c.UserRepository_
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
