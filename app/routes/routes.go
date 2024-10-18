package routes

import (
	"backend-go/app/controllers"
	"backend-go/app/repositories"
	"backend-go/app/services"
	"net/http"

	"github.com/gorilla/mux"
)

// Funcion para inizializar la ruta, recibieindo todo el body con info y leyendo a donde se tieine que ir:
func InitiallizeRoutes(router *mux.Router) {
	repositorie := &repositories.UserRepository{}
	service := &services.UserService{Repo: repositorie}
	controller := &controllers.UserController{Service: service}

	router.HandleFunc("/users/", controller.CreateUser).Methods(http.MethodPost)
	// Lo que va a utilizar . como lo va a utilizar (method)

	// Obtener todos los usuarios
	router.HandleFunc("/users/", controller.GetAllUsers).Methods(http.MethodGet)

	// Obtener usuario por ID
	router.HandleFunc("/users/{id}", controller.GetUserById).Methods(http.MethodGet)

	// Obtener usuario por rol
	router.HandleFunc("/users/rol/{rol}", controller.GetUserByRol).Methods(http.MethodGet)

	// Actualiar usuario
	router.HandleFunc("/users/{id}", controller.UpdateUser).Methods(http.MethodPut)

	// Borrar Usuario:
	router.HandleFunc("/users/{id}", controller.DeleteUser).Methods(http.MethodDelete)
}
