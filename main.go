package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/Antonpractic/go_final_project/database"
	"github.com/Antonpractic/go_final_project/server"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Download error: %v", err)
	}

	database.InitDB()

	server.StartServer()
}
