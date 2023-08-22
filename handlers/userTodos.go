package handlers

import (
	"encoding/json"
	"go-server/utils"
	"net/http"
)

func UserTodos(w http.ResponseWriter, r *http.Request) {

	// Get param for user_id

	user_id := r.URL.Query().Get("user_id")

	if user_id == "" {
		http.Error(w, "No user id provided.", http.StatusBadRequest)
	}

	// Initialise db

	db := utils.GetDB()

	// Query db for user rows

	query := "SELECT * FROM todos WHERE user_id = $1"
	rows, err := db.Query(query, user_id)

	if err != nil {
		http.Error(w, "There's been an error querying the database", http.StatusBadRequest)
		return
	}

	results := []FullTodo{}

	// Scan rows

	for rows.Next() {
		var todo FullTodo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Status, &todo.UserId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		results = append(results, todo)
	}

	// Marshal JSON

	json, err := json.Marshal(results)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Return to FE
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)

}
