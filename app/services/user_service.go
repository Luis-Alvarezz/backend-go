package services

import (
	"backend-go/app/models"
	"backend-go/app/repositories"
	"errors"

	"golang.org/x/crypto/bycrypt"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func (s *UserService) isDuplicated(usuario string) (bool, error) {
	users, err := s.Repo.GetAllUsers()
	if err != nil {
		return false, err
	}

	for _, u := range users {
		if u.Usuario == usuario {
			return true, nil
		}
	}
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
}

// Función para Crear Usuario
func (s *UserService) CreateUser(user models.Usuario) error {
	isDuplicated, err := s.isUserDuplicated(user.Usuario)
	if err != nil {
		return err
	}

	if isDuplicated {
		return errors.New("El usuario ya existe")
	}

	isNameDuplicated, err := s.isNameDuplicated(user.Nombre, user.Apaterno, user.Amaterno)

	if err != nil {
		return err
	}

	if isNameDuplicated {
		return errors.New("El Nombre Completo ya existe")
	}

	hashedPassword, err := bycrypt.GenerateFromPassword([]byte(user.Password), bycrypt.DefaultCost)

	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	user.Imagen = "default.png"

	err = s.Repo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}
