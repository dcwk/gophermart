package application

import (
	"fmt"
	"io"
	"net/http"

	"github.com/dcwk/gophermart/internal/utils/auth"
)

func (app *Application) LoadOrder(w http.ResponseWriter, r *http.Request) {
	orderNumber, err := io.ReadAll(r.Body)
	if err != nil {
		err := fmt.Errorf("couldn't get order number from request")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userID := auth.GetUserIDFromCtx(r.Context())

	err = app.Container.LoadOrderService().Handle(r.Context(), string(orderNumber), userID)
	if err != nil {
		app.Container.Logger().Info(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//if err := json.NewEncoder(w).Encode(resp); err != nil {
	//	app.Container.Logger().Error(err.Error())
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	//w.WriteHeader(http.StatusOK)
	return
}
