package application

import (
	"encoding/json"
	"net/http"

	"github.com/dcwk/gophermart/internal/models"
)

type LoginResponse struct {
	Token string `json:"token"`
}

func (app *Application) Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		app.Container.Logger().Info(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	token, err := app.Container.AuthUserService().Authenticate(r.Context(), &user)

	resp := LoginResponse{Token: token}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		app.Container.Logger().Info(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	return
}
