package application

import (
	"encoding/json"
	"net/http"

	"github.com/dcwk/gophermart/internal/models"
	"github.com/dcwk/gophermart/internal/utils/auth"
)

type OrdersListResponse struct {
	OrdersList []*models.Order `json:"orders"`
}

func (app *Application) GetOrdersList(w http.ResponseWriter, r *http.Request) {
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

	resp := OrdersListResponse{OrdersList: orders}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		app.Container.Logger().Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return
}
