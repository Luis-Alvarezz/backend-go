package repositories

import (
	"backend-go/app/models"
	"context"

	"cloud.google.com/go/firestore"
)

type UserRepository struct {
	Client *firestore.Client
}

func (r *UserRepository) CreateUser(user models.Usuario) error {
	_, _, err := r.Client.Collection("usuario").Add(context.Background(), user) // _, _, Es para
	return err
}
