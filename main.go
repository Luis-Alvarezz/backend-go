package main

// File for initial routes and control of all page

import (
	"backend-go/app/routes"
	"backend-go/config"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Inicializamos FIrebase:
	config.InitialFirebaseApp()
	// Inicializar las rutas:
	router := mux.NewRouter()
	routes.InitiallizeRoutes(router)
	// router.Use(middleware.AuthMiddleware) // Para indicar las rutas
	log.Println("Servidor Trabajando en el puerto: 8080 ")
	log.Fatal(http.ListenAndServe(":8000", router))
}
