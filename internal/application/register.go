package application

import (
	"encoding/json"
	"net/http"
)

type RegisterRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	ID int64 `json:"userId"`
}

func (app *Application) Register(w http.ResponseWriter, r *http.Request) {
	var request RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		app.Container.Logger().Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := app.Container.RegisterUserService().CreateUser(
		r.Context(),
		request.Login,
		request.Password,
	)
	if err != nil {
		app.Container.Logger().Info(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := RegisterResponse{
		ID: user.ID,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		app.Container.Logger().Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return
}
