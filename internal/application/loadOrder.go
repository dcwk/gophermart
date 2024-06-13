package application

import (
	"fmt"
	"io"
	"net/http"

	"github.com/dcwk/gophermart/internal/models"
	"github.com/dcwk/gophermart/internal/usecase"
	"github.com/dcwk/gophermart/internal/utils/auth"
)

func (app *Application) LoadOrder(w http.ResponseWriter, r *http.Request) {
	orderNumber, err := io.ReadAll(r.Body)
	resChannel := make(chan string)
	errChannel := make(chan error)
	orderChannel := make(chan models.AccrualOrder)
	if err != nil {
		err := fmt.Errorf("couldn't get order number from request")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userID := auth.GetUserIDFromCtx(r.Context())

	go app.Container.LoadOrderHandler().Handle(r.Context(), orderChannel, string(orderNumber))
	go func() {
		code, err := app.Container.CreateOrderHandler().Handle(r.Context(), orderChannel, string(orderNumber), userID)
		resChannel <- code
		errChannel <- err
	}()

	code := <-resChannel
	err = <-errChannel
	if code == "" && err == nil {
		w.WriteHeader(http.StatusAccepted)
		return
	}
	if err != nil {
		app.Container.Logger().Error(err.Error())
	}

	switch code {
	case usecase.InvalidOrder:
		w.WriteHeader(http.StatusAccepted)
	case usecase.OrderAlreadyExists:
		w.WriteHeader(http.StatusOK)
	case usecase.ForbiddenOrder:
		w.WriteHeader(http.StatusConflict)
	case usecase.IncorrectOrderNumber:
		w.WriteHeader(http.StatusUnprocessableEntity)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}
