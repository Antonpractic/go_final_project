package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Antonpractic/go_final_project/database"
	_ "modernc.org/sqlite"
)

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
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

	_, err = database.DoneTask(h.DB, id)
	if err != nil {
		http.Error(w, `{"Не получилось удалить задачу"}`, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json, charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(fmt.Sprint(`{}`)))
	return

}
