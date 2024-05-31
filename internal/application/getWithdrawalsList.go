package application

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dcwk/gophermart/internal/utils/auth"
)

type withdrawResponse struct {
	Order       string    `json:"order"`
	Sum         float64   `json:"sum"`
	ProcessedAt time.Time `json:"processed_at"`
}

func (app *Application) GetWithdrawalsList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	userID := auth.GetUserIDFromCtx(r.Context())
	withdrawals, err := app.Container.GetWithdrawalsHandler().Handle(
		r.Context(),
		userID,
	)
	if err != nil {
		app.Container.Logger().Info(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := make([]*withdrawResponse, len(withdrawals))
	for i, withdrawal := range withdrawals {
		resp[i] = &withdrawResponse{
			Order:       withdrawal.OrderNumber,
			Sum:         withdrawal.Value,
			ProcessedAt: withdrawal.CreatedAt,
		}
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		app.Container.Logger().Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
