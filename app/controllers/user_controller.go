package controllers

import (
	"backend-go/app/models"
	"backend-go/app/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type UserController struct {
	Service *services.UserService
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.Usuario
	json.NewDecoder(r.Body).Decode(&user)
	err := c.Service.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// Funcion para obtener todos los usuarios:
func (c *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.Service.Repo.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

// Obtener por ID:
func (c *UserController) GetUserById(w http.ResponseWriter, r *http.Request) { // Request que llega y response que realizará
	id := mux.Vars(r)["id"]
	user, err := c.Service.GetUserById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// Obtener usuario por Rol:
func (c *UserController) GetUserByRol(resWritt http.ResponseWriter, request *http.Request) {
	rol := mux.Vars(request)["rol"]
	user, err := c.Service.Repo.GetUserByRol(rol)
	if err != nil {
		http.Error(resWritt, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(resWritt).Encode(user)
}

// Actualizar Usuario
func (c *UserController) UpdateUser(resWrti http.ResponseWriter, requers *http.Request) {
	id := mux.Vars(requers)["id"]
	var user models.Usuario
	if err := json.NewDecoder(requers.Body).Decode(&user); err != nil {
		http.Error(resWrti, err.Error(), http.StatusBadRequest)
		return
	}
	if err := c.Service.UpdateUser(id, user); err != nil {
		http.Error(resWrti, err.Error(), http.StatusInternalServerError)
		return
	}
	resWrti.WriteHeader(http.StatusOK)
}

// Borrar Usuario:
func (c *UserController) DeleteUser(resWrti http.ResponseWriter, request *http.Request) {
	id := mux.Vars(request)["id"]
	// user, err := c.Service.DeleteUser(id)
	if err := c.Service.DeleteUser(id); err != nil {
		http.Error(resWrti, err.Error(), http.StatusInternalServerError)
		return
	}
	resWrti.WriteHeader(http.StatusOK)
}
