package server

import (
	"log"
	"net/http"
	"os"

	"github.com/Antonpractic/go_final_project/constants"
	"github.com/Antonpractic/go_final_project/database"
	"github.com/Antonpractic/go_final_project/handlers"
	"github.com/go-chi/chi"
)

func StartServer() {

	Handler := handlers.NewHandler(database.Db)

	r := chi.NewRouter()

	fs := http.FileServer(http.Dir(constants.WebDir))

	r.Get("/api/task", Handler.GetTaskId)
	r.Get("/api/nextdate", Handler.GetDateTask)
	r.Get("/api/tasks", Handler.GetTasks)
	r.Post("/api/task", Handler.PostTask)
	r.Post("/api/task/done", Handler.DoneTask)
	r.Put("/api/task", Handler.PutTask)
	r.Delete("/api/task", Handler.DeleteTask)

	r.Handle("/*", fs)

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
