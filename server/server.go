package server

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func StartServer() {

	r := chi.NewRouter()

	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}
	log.Println("Port 7540")
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Server start error: %s\n", err.Error())
		return
	}
}
