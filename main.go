package main

import (
	"database/sql"
	"go-server/handlers"
	"go-server/middleware"
	"go-server/utils"
	"net/http"
)

type Service struct {
	db *sql.DB
}

func main() {

	utils.Initialise()

	rate := 1
	limiter := middleware.NewRateLimiter(rate)

	mux := http.NewServeMux()

	wrappedHandler := middleware.Auth(limiter.Limit(mux))

	mux.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handlers.CreateTodo(w, r)
		} else if r.Method == http.MethodGet {
			handlers.SingleTodo(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/all-todos", handlers.AllTodos)

	mux.HandleFunc("/user-todos", handlers.UserTodos)

	mux.HandleFunc("/delete-todo", handlers.DeleteTodo)

	mux.HandleFunc("/update-todo", handlers.UpdateTodo)

	http.ListenAndServe(":8080", wrappedHandler)

	defer utils.GetDB().Close()

}
