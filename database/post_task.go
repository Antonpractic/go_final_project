package database

import (
	"database/sql"
	"log"

	"github.com/Antonpractic/go_final_project/models"
	_ "modernc.org/sqlite"
)

func AddTask(db *sql.DB, task models.Task) (int64, error) {
	res, err := db.Exec("INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?)",
		task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		log.Printf("Не удалось добавить задачу с заголовком: %v", task.Title)
		return 0, err
	}

	Id, err := res.LastInsertId()
	if err != nil {
		log.Println("Не удалось получить ID добавленной задачи", err)
		return 0, err
	}
	return Id, nil
}
