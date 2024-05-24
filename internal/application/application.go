package application

import (
	"net/http"

	"github.com/dcwk/gophermart/internal/config"
	"github.com/dcwk/gophermart/internal/container"
	"github.com/dcwk/gophermart/internal/utils/middleware"
	"github.com/go-chi/chi/v5"
)

type Application struct {
	Container *container.Container
}

func Run(conf *config.ServerConf) {
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

	r.Post("/api/user/register", app.Register)
	r.Post("/api/user/login", app.Login)
	r.Post("/api/user/orders", app.LoadOrder)
	r.Get("/api/user/orders", app.GetOrdersList)
	r.Get("/api/user/balance", app.GetUserBalance)
	r.Post("/api/user/balance/withdraw", app.WithdrawRequest)
	r.Get("/api/user/withdrawals", app.GetWithdrawalsList)

	return r
}
