package application

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dcwk/gophermart/internal/utils/auth"
)

type orderResponse struct {
	Number     string    `json:"number"`
	Status     string    `json:"status"`
	Accrual    float64   `json:"accrual"`
	UploadedAt time.Time `json:"uploaded_at"`
}

func (app *Application) GetOrdersList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	userID := auth.GetUserIDFromCtx(r.Context())
	orders, err := app.Container.GetOrdersService().Handle(
		r.Context(),
		userID,
	)
	if err != nil {
		app.Container.Logger().Info(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := make([]*orderResponse, len(orders))
	for i, order := range orders {
		resp[i] = &orderResponse{
			Number:     order.Number,
			Status:     order.Status,
			Accrual:    order.Accrual,
			UploadedAt: order.CreatedAt,
		}
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		app.Container.Logger().Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
