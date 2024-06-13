package application

import (
	"encoding/json"
	"net/http"
)

type RegisterRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (app *Application) Register(w http.ResponseWriter, r *http.Request) {
	var request RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		app.Container.Logger().Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := app.Container.RegisterUserHandler().Handle(
		r.Context(),
		request.Login,
		request.Password,
	)
	if err != nil {
		app.Container.Logger().Info(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Authorization", "Bearer "+token)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
