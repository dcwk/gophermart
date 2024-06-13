package application

import (
	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (app *Application) Login(w http.ResponseWriter, r *http.Request) {
	var request LoginRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		app.Container.Logger().Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := app.Container.AuthUserHandler().Handle(
		r.Context(),
		request.Login,
		request.Password,
	)
	if err != nil {
		app.Container.Logger().Info(err.Error())
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	w.Header().Set("Authorization", "Bearer "+token)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
