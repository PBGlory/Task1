package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var task string

func GetHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Получен GET-запрос")
	var tasks []Task

	if err := DB.Find(&tasks).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Получен POST-запрос")

	var requestBody struct {
		Task   string `json:"task"`
		IsDone bool   `json:"is_Done"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if requestBody.Task == "" {
		http.Error(w, "Поле task обязательно", http.StatusBadRequest)
		return
	}

	task := Task{
		Task:   requestBody.Task,
		IsDone: requestBody.IsDone,
	}

	if err := DB.Create(&task).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := map[string]string{
		"message": fmt.Sprintf("Задача успешно создана: %s", task.Task),
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	InitDB()

	DB.AutoMigrate(&Task{})

	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", GetHandler).Methods("GET")
	router.HandleFunc("/api/tasks", PostHandler).Methods("POST")

	fmt.Println("Сервер запущен по адресу: http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
