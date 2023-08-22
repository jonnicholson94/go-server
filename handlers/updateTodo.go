package handlers

import (
	"encoding/json"
	"fmt"
	"go-server/utils"
	"net/http"
)

type UpdatedStatus struct {
	Status bool `json:"status"`
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {

	// Get id of row to update
	id := r.URL.Query().Get("id")
	var status UpdatedStatus

	err := json.NewDecoder(r.Body).Decode(&status)

	if id == "" {
		http.Error(w, "No id provided", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, "No status provided", http.StatusBadRequest)
	}

	st := status.Status

	// Initialise db

	db := utils.GetDB()

	// Write query and prep statement

	query := "UPDATE todos SET status = $1 WHERE id = $2"

	stmt, err := db.Prepare(query)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute query

	data, err := stmt.Exec(st, id)

	// Check for error, otherwise return

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(data)
}
