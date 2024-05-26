package application

import (
	"fmt"
	"net/http"

	"github.com/dcwk/gophermart/internal/utils/auth"
)

func (app *Application) GetOrdersList(w http.ResponseWriter, r *http.Request) {
	app.Container.Logger().Info(fmt.Sprintf("User with id: %v", auth.GetUserIDFromCtx(r.Context())))
}
