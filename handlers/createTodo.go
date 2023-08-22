package handlers

import (
	"encoding/json"
	"fmt"
	"go-server/utils"
	"net/http"
)

func CreateTodo(w http.ResponseWriter, r *http.Request) {

	var todo NewTodo

	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		http.Error(w, "Incorrect body format provided", http.StatusBadRequest)
	}

	fmt.Println("Received data", todo)

	db := utils.GetDB()

	query := "INSERT INTO todos (title, status, user_id) VALUES ($1, $2, $3)"

	data, err := db.Exec(query, todo.Title, todo.Status, todo.UserId)

	if err != nil {
		http.Error(w, "Failed to add todo to database", http.StatusBadRequest)
	}

	fmt.Println(data)

}
