package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"

	"github.com/Antonpractic/go_final_project/models"
)

func PutTask(db *sql.DB, task models.Task) (int64, error) {

	res, err := db.Exec("UPDATE scheduler SET date = ?, title = ?, comment = ?, repeat = ? WHERE id = ?",
		task.Date, task.Title, task.Comment, task.Repeat, task.ID)
	if err != nil {
		log.Printf("Не удалось обновить задачу с ID: %v", task.ID)
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
