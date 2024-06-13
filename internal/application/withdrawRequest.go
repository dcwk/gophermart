package application

import (
	"encoding/json"
	"net/http"

	"github.com/dcwk/gophermart/internal/usecase"
	"github.com/dcwk/gophermart/internal/utils/auth"
)

type withdrawRequest struct {
	OrderNumber string  `json:"order"`
	Sum         float64 `json:"sum"`
}

func (app *Application) WithdrawRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	var request withdrawRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		app.Container.Logger().Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID := auth.GetUserIDFromCtx(r.Context())
	code, err := app.Container.WithdrawRequestHandler().Handle(
		r.Context(),
		userID,
		request.OrderNumber,
		request.Sum,
	)
	if code == "" && err == nil {
		w.WriteHeader(http.StatusOK)
		return
	}

	switch code {
	case usecase.NotEnoughPoints:
		w.WriteHeader(http.StatusPaymentRequired)
	case usecase.NotFound:
		w.WriteHeader(http.StatusNotFound)
	case usecase.ForbiddenOrder:
		w.WriteHeader(http.StatusConflict)
	case usecase.IncorrectOrderNumber:
		w.WriteHeader(http.StatusUnprocessableEntity)
	default:
		if err != nil {
			app.Container.Logger().Error(err.Error())
		}
		w.WriteHeader(http.StatusInternalServerError)
	}
}
