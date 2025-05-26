package handlers

import (
	"fmt"

	"net/http"

	_ "modernc.org/sqlite"

	"github.com/Antonpractic/go_final_project/database"
	"github.com/Antonpractic/go_final_project/helpers"
	"github.com/Antonpractic/go_final_project/models"
)

func (h *Handler) PostTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := helpers.DecodeJSON(r.Body, &task); err != nil {
		http.Error(w, `{"Не удалось расшифровать JSON"}`, http.StatusBadRequest)
		return
	}

	if err := helpers.CheckTask(&task); err != nil {
		http.Error(w, fmt.Sprintf(`{"Ошибка": "%s"}`, err.Error()), http.StatusBadRequest)
		return
	}

	id, err := database.AddTask(h.DB, task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json, charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte(fmt.Sprintf(`{"ID": %d}`, id)))
}
