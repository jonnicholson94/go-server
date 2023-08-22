package handlers

import (
	"encoding/json"
	"go-server/utils"
	"net/http"
)

func AllTodos(w http.ResponseWriter, r *http.Request) {

	db := utils.GetDB()

	query := "SELECT * FROM todos"
	rows, err := db.Query(query)

	if err != nil {
		http.Error(w, "There's been an error fetching the data", http.StatusBadRequest)
	}

	todos := []FullTodo{}

	for rows.Next() {
		var todo FullTodo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Status, &todo.UserId)
		if err != nil {
			http.Error(w, "There's been an error processing the data", http.StatusBadRequest)
		}
		todos = append(todos, todo)
	}

	json, err := json.Marshal(todos)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)

}
