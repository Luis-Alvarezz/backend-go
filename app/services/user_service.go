package services

import (
	"backend-go/app/models"
	"backend-go/app/repositories"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func (s *UserService) isUserDuplicated(usuario string) (bool, error) {
	users, err := s.Repo.GetAllUsers()
	if err != nil {
		return false, err
	}

	for _, u := range users {
		if u.Usuario == usuario {
			return true, nil
		}
	}
	return false, nil
}

func (s *UserService) isNameDuplicated(nombre, apaterno, amaterno string) (bool, error) {
	users, err := s.Repo.GetAllUsers()
	if err != nil {
		return false, err
	}

	for _, u := range users {
		if u.Nombre == nombre && u.Apaterno == apaterno && u.Amaterno == amaterno {
			return true, nil
		}
	}
	return false, nil
}

// Función para Crear Usuario
func (s *UserService) CreateUser(user models.Usuario) error {
	isDuplicated, err := s.Repo.GetUserByUsername(user.Usuario)
	if err != nil {
		return err
	}

	if isDuplicated != nil {
		return errors.New("El usuario ya existe")
	}

	isNameDuplicated, err := s.isNameDuplicated(user.Nombre, user.Apaterno, user.Amaterno)

	if err != nil {
		return err
	}

	if isNameDuplicated {
		return errors.New("El Nombre Completo ya existe")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// Almacenamos ¡password hasheado
	user.Password = string(hashedPassword)

	// user.Imagen = "default.png"
	return s.Repo.CreateUser(user)
}

func (s *UserService) UpdateUser(id string, user models.Usuario) error {
	userID, err := s.Repo.GetUserById(id)
	if err != nil {
		return err
	}
	if userID == nil {
		return errors.New("Usuario NO encontrado")
	}
	userID.Nombre = user.Nombre
	userID.Apaterno = user.Apaterno
	userID.Amaterno = user.Amaterno
	userID.Direccion = user.Direccion
	userID.Telefono = user.Telefono
	userID.Ciudad = user.Ciudad
	userID.Estado = user.Estado
	userID.Rol = user.Rol
	userID.Imagen = user.Imagen

	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil
		}
		userID.Password = string(hashedPassword)
	}
	return s.Repo.UpdateUser(id, *userID)
}

/* Funcion para eliminar usuario */
func (s *UserService) DeleteUser(id string) error {
	userID, err := s.Repo.GetUserById(id)
	if err != nil {
		return nil
	}
	if userID == nil {
		return errors.New("Usuario NO encontrado")
	}
	return s.Repo.DeleteUser(id)
}

func (s *UserService) GetUserById(id string) (*models.Usuario, error) {
	return s.Repo.GetUserById(id)
}

func (s *UserService) GetUserByUsername(username string) (*models.Usuario, error) {
	return s.Repo.GetUserByUsername(username)
}

func (s *UserService) GetUserByRol(rol string) ([]models.Usuario, error) {
	return s.Repo.GetUserByRol(rol)
}
