package application

import (
	"encoding/json"
	"net/http"

	"github.com/dcwk/gophermart/internal/utils/auth"
)

type UserBalanceResponse struct {
	Accrual    float64 `json:"accrual"`
	Withdrawal float64 `json:"withdrawal"`
}

func (app *Application) GetUserBalance(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromCtx(r.Context())
	userBalance, err := app.Container.GetUserBalanceService().Handle(
		r.Context(),
		userID,
	)
	if err != nil {
		app.Container.Logger().Info(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := UserBalanceResponse{Accrual: userBalance.Accrual, Withdrawal: userBalance.Withdrawal}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		app.Container.Logger().Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return

}
