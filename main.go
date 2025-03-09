package main

import (
	"Task1/internal/database"
	"Task1/internal/handlers"
	"Task1/internal/taskService"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	database.InitDB()

	database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewTaskService(repo)

	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", handler.GetTaskHandler).Methods("GET")
	router.HandleFunc("/api/tasks", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/tasks/{id}", handler.UpdateTaskHandler).Methods("PATCH")
	router.HandleFunc("/api/tasks/{id}", handler.DeleteTaskHandler).Methods("DELETE")

	fmt.Println("Сервер запущен по адресу: http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
