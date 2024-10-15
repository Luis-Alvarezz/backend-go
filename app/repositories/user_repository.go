package repositories

import (
	"backend-go/app/models"
	"backend-go/config"
	"context"

	"google.golang.org/api/iterator"
)

type UserRepository struct {
}

func (r *UserRepository) CreateUser(user models.Usuario) error {
	contexto := context.Background()
	client, err := config.FirebaseApp.Firestore(contexto)
	if err != nil {
		return err
	}
	defer client.Close() // Similar al Await, espera ejecucion de funcion y se cierra

	_, _, err = client.Collection("usuarios_lenguajes").Add(contexto, user)
	return err
}

func (r *UserRepository) GetAllUsers() ([]models.Usuario, error) {
	ctx := context.Background()
	client, err := config.FirebaseApp.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	var users []models.Usuario
	iter := client.Collection("usuarios_lenguajes").Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var user models.Usuario
		doc.DataTo(&user)
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) GetUserById(id string) (*models.Usuario, error) {
	contexto := context.Background()
	client, err := config.FirebaseApp.Firestore(contexto)
	if err != nil {
		return nil, err
	}

	defer client.Close()
	doc, err := client.Collection("usuarios_lenguajes").Doc(id).Get(contexto)
	if err != nil {
		return nil, err
	}
	var user_id models.Usuario
	doc.DataTo(&user_id)
	user_id.ID = doc.Ref.ID
	return &user_id, nil
}

func (r *UserRepository) GetUserByUsername(username string) (*models.Usuario, error) { // username -> Debe estar igual que en la conleccion de DB
	contx := context.Background()
	client, err := config.FirebaseApp.Firestore(contx)
	if err != nil {
		return nil, err
	}

	defer client.Close()
	iter := client.Collection("usuarios_lenguajes").Where("usuario", "==", username).Documents(contx)
	doc, err := iter.Next()

	if err == iterator.Done {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	var usernameColl models.Usuario
	doc.DataTo(&usernameColl)
	usernameColl.ID = doc.Ref.ID
	return &usernameColl, nil
}

func (r *UserRepository) GetUserByRol(rol string) ([]models.Usuario, error) {
	contexto := context.Background()
	client, err := config.FirebaseApp.Firestore(contexto)
	if err != nil {
		return nil, err
	}

	var users []models.Usuario
	iter := client.Collection("usuarios_lenguajes").Where("rol", "==", rol).Documents(contexto)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var user models.Usuario
		doc.DataTo(&user)
		users = append(users, user)

	}
	return users, nil
}

// Funcion para actualizar a un usuario
func (r *UserRepository) UpdateUser(id string, user models.Usuario) error {
	ctx := context.Background()
	// Conexion:
	client, err := config.FirebaseApp.Firestore(ctx)
	if err != nil {
		return err
	}

	defer client.Close()
	_, err = client.Collection("usuarios_lenguajes").Doc(id).Set(ctx, user)
	return err
}

// Funcion para actualizar a un usuario
func (r *UserRepository) DeleteUser(id string) error {
	ctx := context.Background()
	// Conexion:
	client, err := config.FirebaseApp.Firestore(ctx)
	if err != nil {
		return err
	}

	defer client.Close()
	_, err = client.Collection("usuarios_lenguajes").Doc(id).Delete(ctx)
	return err
}
