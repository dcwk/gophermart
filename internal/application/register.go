package application

import (
	"encoding/json"
	"net/http"

	"github.com/dcwk/gophermart/internal/models"
)

func (app *Application) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		app.Container.Logger().Info(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	newUser, err := app.Container.RegisterUserService().CreateUser(r.Context(), &user)
	if err != nil {
		app.Container.Logger().Info(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(newUser); err != nil {
		app.Container.Logger().Info(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	return
}
