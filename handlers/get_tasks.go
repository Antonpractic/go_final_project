package handlers

import (
	"log"
	"net/http"

	_ "modernc.org/sqlite"

	"github.com/Antonpractic/go_final_project/database"
	"github.com/Antonpractic/go_final_project/helpers"
	"github.com/Antonpractic/go_final_project/models"
)

func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	tasks, err := database.GetAllTasks(h.DB)
	if err != nil {
		http.Error(w, `{"error": "Не удалось получить задачи"}`, http.StatusInternalServerError)
		log.Println("Ошибка при извлечении задач:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json, charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	response := models.TaskResponse{
		Tasks: tasks,
	}

	helpers.EncodeJSON(w, response)

}
