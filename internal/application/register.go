package application

import (
	"net/http"

	"github.com/dcwk/gophermart/internal/models"
)

// TODO: Доработать хэндлер
func (app *Application) Register(w http.ResponseWriter, r *http.Request) {
	userData := models.User{
		Login:    "test",
		Password: "test",
	}
	_, err := app.Container.RegisterUserService().CreateUser(r.Context(), &userData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}
