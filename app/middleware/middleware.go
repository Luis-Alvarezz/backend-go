package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resWrit http.ResponseWriter, request *http.Request) {
		authHandle := request.Header.Get("Authorization")
		if authHandle == "" {
			http.Error(resWrit, "Token no proporcionado", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.Replace(authHandle, "Bearer", "", 1)
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrAbortHandler
			}
			return []byte("PALABRA_SECRETA"), nil
		})
		if err != nil || !token.Valid {
			http.Error(resWrit, "Token Inválido", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(resWrit, request)
	})
}
