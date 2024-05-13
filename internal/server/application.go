package server

import (
	"net/http"

	"github.com/dcwk/gophermart/internal/config"
	"github.com/dcwk/gophermart/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func Run(conf *config.ServerConf) {
	if err := http.ListenAndServe(conf.RunAddress, Router()); err != nil {
		panic(err)
	}
}

func Router() chi.Router {
	r := chi.NewRouter()

	h := handlers.Handlers{}

	r.Post("/api/user/register", h.Register)
	r.Post("/api/user/login", h.Login)
	r.Post("/api/user/orders", h.LoadOrder)
	r.Get("/api/user/orders", h.GetOrdersList)
	r.Get("/api/user/balance", h.GetUserBalance)
	r.Post("/api/user/balance/withdraw", h.WithdrawRequest)
	r.Get("/api/user/withdrawals", h.GetWithdrawalsList)

	return r
}