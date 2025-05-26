package handlers

import (
	"fmt"
	"net/http"
	"time"

	_ "modernc.org/sqlite"

	"github.com/Antonpractic/go_final_project/constants"
	"github.com/Antonpractic/go_final_project/helpers"
)

func (h *Handler) GetDateTask(w http.ResponseWriter, r *http.Request) {
	nowstr := r.URL.Query().Get("now")
	if nowstr == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"Сейчас дата отсутствует"}`))
		return
	}

	now, err := time.Parse(constants.FormatDate, nowstr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"Неверный формат даты"}`))
		return
	}

	date := r.URL.Query().Get("date")
	if date == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"Дата отсутствует"}`))
		return
	}

	repeat := r.URL.Query().Get("repeat")
	if repeat == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"Повторное отсутствие"}`))
		return
	}

	nextDate, err := helpers.NextDate(now, date, repeat)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	w.Header().Set("Content-Type", "application/json, charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(fmt.Sprintf("%s", nextDate)))
}
