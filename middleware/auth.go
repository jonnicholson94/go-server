package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func Auth(next http.Handler) http.Handler {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Env variables not loaded correctly")
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		token := os.Getenv("TOKEN")

		if authHeader != token {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
