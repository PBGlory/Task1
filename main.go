package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	InitDB()

	DB.AutoMigrate(&Task{})

	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", GetHandler).Methods("GET")
	router.HandleFunc("/api/tasks", PostHandler).Methods("POST")
	router.HandleFunc("/api/tasks/{id}", PatchHandler).Methods("PATCH")
	router.HandleFunc("/api/tasks/{id}", DeleteHandler).Methods("DELETE")

	fmt.Println("Сервер запущен по адресу: http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
