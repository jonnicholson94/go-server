package handlers

import (
	"go-server/utils"
	"net/http"
)

func DeleteTodo(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "No id was provided, so the todo can't be deleted", http.StatusBadRequest)
		return
	}

	// Initialise database variable

	db := utils.GetDB()

	query := "DELETE FROM todos WHERE id = $1"

	stmt, err := db.Prepare(query)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = stmt.Exec(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
