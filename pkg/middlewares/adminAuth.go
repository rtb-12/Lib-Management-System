package middlewares

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func AdminAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		cookie, err := request.Cookie("jwt")
		if err != nil {
			http.Error(writer, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenStr := cookie.Value
		claims := &jwt.MapClaims{}
		_, err = jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			http.Error(writer, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(writer, request)
	})
}
