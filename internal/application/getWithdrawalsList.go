package application

import (
	"encoding/json"
	"net/http"

	"github.com/dcwk/gophermart/internal/models"
	"github.com/dcwk/gophermart/internal/utils/auth"
)

type WithdrawalsListResponse struct {
	WithdrawalsList []*models.Withdrawal `json:"withdrawals"`
}

func (app *Application) GetWithdrawalsList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	userID := auth.GetUserIDFromCtx(r.Context())
	withdrawals, err := app.Container.GetWithdrawalsService().Handle(
		r.Context(),
		userID,
	)
	if err != nil {
		app.Container.Logger().Info(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := WithdrawalsListResponse{WithdrawalsList: withdrawals}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		app.Container.Logger().Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
