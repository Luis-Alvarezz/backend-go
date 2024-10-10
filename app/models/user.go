package models

// Definimos la estructura que vamos a manejar:
type Usuario struct {
	ID        string `json:"id"`
	Nombre    string `json:"nombre"`
	Apaterno  string `json:"apaterno"`
	Amaterno  string `json:"amaterno"`
	Direccion string `json:"direccion"`
	Telefono  string `json:"telefono"`
	Ciudad    string `json:"ciudad"`
	Estado    string `json:"estado"`
	Usuario   string `json:"usuario"`
	Password  string `json:"-"`
	Rol       string `json:"rol"`
	Imagen    string `json:"imagen"`
}
