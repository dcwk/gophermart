package application

import (
	"net/http"

	"github.com/dcwk/gophermart/internal/config"
	"github.com/dcwk/gophermart/internal/container"
	"github.com/dcwk/gophermart/internal/utils/middleware"
	"github.com/dcwk/gophermart/migrations"
	"github.com/go-chi/chi/v5"
)

type Application struct {
	Container *container.Container
}

func Run(conf *config.ServerConf) {
	err := migrations.RunMigrations(conf.DatabaseDSN)
	if err != nil {
		panic(err)
	}

	app := Application{
		Container: container.NewContainer(conf),
	}
	if err := http.ListenAndServe(conf.RunAddress, app.Router()); err != nil {
		panic(err)
	}
}

func (app *Application) Router() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.RequestMiddleware(app.Container.Logger()))

	r.Route("/api/user", func(r chi.Router) {
		r.Post("/register", app.Register)
		r.Post("/login", app.Login)

		r.With(middleware.JwtAuth).Get("/orders", app.GetOrdersList)
		r.With(middleware.JwtAuth).Get("/balance", app.GetUserBalance)
		r.With(middleware.JwtAuth).Get("/withdrawals", app.GetWithdrawalsList)

		r.With(middleware.JwtAuth).Post("/orders", app.LoadOrder)
		r.With(middleware.JwtAuth).Post("/balance/withdraw", app.WithdrawRequest)
	})

	return r
}
