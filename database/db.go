package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"

	"github.com/Antonpractic/go_final_project/constants"
)

var Db *sql.DB

func InitDB() {

	dbFile := os.Getenv("TODO_DBFILE")
	if dbFile == "" {
		dbFile = filepath.Join("./scheduler.db")
	}

	_, err := os.Stat(dbFile)
	var install bool
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("Файл c данными не найден, создан новый файл:", dbFile)
			install = true
		} else {
			log.Fatal("Ошибка при проверке файла с данными:", err)
		}
	}

	Db, err = sql.Open("sqlite", dbFile)
	if err != nil {
		log.Fatal("Ошибка при открытии данных:", err)
	}

	if install {
		fileContent, err := os.ReadFile(constants.CreateTable)
		if err != nil {
			log.Fatal("Ошибка при чтении table.sql", err)
		}

		_, err = Db.Exec(string(fileContent))
		log.Println("Таблица успешно создана")
		if err != nil {
			log.Fatal("Ошибка при выполнении table.sql", err)
		}
	}
}
