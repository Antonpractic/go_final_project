package helpers

import (
	"fmt"
	"time"

	"github.com/Antonpractic/go_final_project/constants"
	"github.com/Antonpractic/go_final_project/models"
	_ "modernc.org/sqlite"
)

func CheckTask(task *models.Task) error {

	if task.Title == "" {
		return fmt.Errorf("отсутствует заголовок")
	}

	if task.Date == "" {
		task.Date = time.Now().Format(constants.FormatDate)
	}

	_, err := time.Parse(constants.FormatDate, task.Date)
	if err != nil {
		return fmt.Errorf("неверный формат даты, ожидаемый ГГГГММДД")
	}

	if task.Date < time.Now().Format(constants.FormatDate) {
		if task.Repeat != "" {
			nextDate, err := NextDate(time.Now(), task.Date, task.Repeat)
			if err != nil {
				return fmt.Errorf("не удалось рассчитать следующую дату")
			}
			task.Date = nextDate
		} else {
			task.Date = time.Now().Format(constants.FormatDate)
		}
	}

	return nil
}
