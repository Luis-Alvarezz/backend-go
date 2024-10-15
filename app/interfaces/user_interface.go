package interfaces

// Prototipos o Funciones que se van a crear sobre cada modelo en el repositorio

import "backend-go/app/models"

type IUser interface {
	// Funciones que vamos a utilizar
	CreateUser(user models.Usuario) error
	GetAllUsers() ([]models.Usuario, error) // El segundo Parentesis es para retornar Arreglo de Objetos
	GetUserById(id string) (*models.Usuario, error)
	GetUserByUsername(usermane string) (*models.Usuario, error) // Retorna el apuntador
	GetUserByRol(rol string) ([]models.Usuario, error)
	UpdateUser(id string, user models.Usuario) error
	DeleteUser(id string) error
}
