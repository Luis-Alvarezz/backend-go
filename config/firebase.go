package config

// Importacion de Librerias a Utilizar:
import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

/* funct Hello (name string) string */
/* funct FunctionName (name Parameter Type) Return Type */
func InitialFirebaseApp() *firebase.App {
	opt := option.WithCredentialsFile("./firebaseServiceAccount.json") // La ruta de Service Account
	app, err := firebase.NewApp(context.Background(), nil, opt)        // Esto es Try-Catch
	if err != nil {
		log.Fatalf("Error al inicializar Firabase App: %v", err)
	}
	return app
}

func GetAuthClient(app *firebase.App) *auth.Client {
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Error al obtener el Cliente: %w", err)
	}
	return client
}
