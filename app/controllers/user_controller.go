package controllers

import (
	"backend-go/app/models"
	"backend-go/app/service"
	"encoding/json"
	"net/http"
)

type UserController struct {
	Service *service.UserService
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.Usuario
	json.NewDecoder(r.Body).Decode(&user)
	err := c.Service.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternaServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
