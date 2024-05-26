package application

import (
	"fmt"
	"net/http"

	"github.com/dcwk/gophermart/internal/models"
	"github.com/dcwk/gophermart/internal/utils/auth"
)

type OrdersListResponse struct {
	OrdersList []models.Order
}

func (app *Application) GetOrdersList(w http.ResponseWriter, r *http.Request) {
	app.Container.Logger().Info(fmt.Sprintf("User with id: %v", auth.GetUserIDFromCtx(r.Context())))
}
