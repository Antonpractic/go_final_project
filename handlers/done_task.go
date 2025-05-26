package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "modernc.org/sqlite"

	"github.com/Antonpractic/go_final_project/database"
	"github.com/Antonpractic/go_final_project/helpers"
)

func (h *Handler) DoneTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"ID не найдено"}`))
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"Неверный формат ID"}`))
		return
	}
	taskFromDB, err := database.GetTaskById(h.DB, id)
	if err != nil {
		http.Error(w, `{"Ошибка базы данных"}`, http.StatusInternalServerError)
		return
	}
	if taskFromDB == nil {
		http.Error(w, `{"Задача не найдена"}`, http.StatusNotFound)
		return
	}

	task := *taskFromDB
	if task.Repeat == "" {
		doneTask, err := database.DoneTask(h.DB, id)
		if err != nil {
			http.Error(w, `{"Не получилось удалить задачу"}`, http.StatusBadRequest)
			return
		}
		log.Println("Задача выполнена", doneTask)

		w.Header().Set("Content-Type", "application/json, charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(fmt.Sprint(`{}`)))
		return
	}

	nextDate, err := helpers.NextDate(time.Now(), task.Date, task.Repeat)
	if err != nil {
		http.Error(w, `{"Не удалось рассчитать следующую дату"}`, http.StatusBadRequest)
		return
	}
	task.Date = nextDate

	updatedTaskID, err := database.PutTask(h.DB, task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("Задача обновлена", updatedTaskID)

	w.Header().Set("Content-Type", "application/json, charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte(fmt.Sprint(`{}`)))

}
