package handlers

import (
	"encoding/json"
	"net/http"

	_ "modernc.org/sqlite"

	"github.com/Antonpractic/go_final_project/database"
	"github.com/Antonpractic/go_final_project/helpers"
	"github.com/Antonpractic/go_final_project/models"
)

func (h *Handler) PostTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := helpers.DecodeJSON(r.Body, &task); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": "Не удалось расшифровать JSON",
		})
		return
	}

	if err := helpers.CheckTask(&task); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	id, err := database.AddTask(h.DB, task)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]int64{
		"id": id,
	})
}
