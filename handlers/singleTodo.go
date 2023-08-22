package handlers

import (
	"encoding/json"
	"go-server/utils"
	"net/http"
)

type FullTodo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"boolean"`
	UserId string `json:"user_id"`
}

func SingleTodo(w http.ResponseWriter, r *http.Request) {

	// Read request param (?)

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "No id provided, so can't return a todo", http.StatusBadRequest)
		return
	}

	// Initialise db

	db := utils.GetDB()

	// Run query

	query := "SELECT * FROM todos WHERE id = $1"
	row := db.QueryRow(query, id)

	var todo FullTodo

	row.Scan(&todo.ID, &todo.Title, &todo.Status, &todo.UserId)

	// Write data back to fe

	json, err := json.Marshal(todo)

	if err != nil {
		http.Error(w, "Failed to write todos", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)

}
