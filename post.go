package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Получен POST-запрос")

	var newTask Task

	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	if newTask.Task == "" {
		http.Error(w, "Поле 'task' обязательно", http.StatusBadRequest)
		return
	}

	if err := DB.Create(&newTask).Error; err != nil {
		http.Error(w, "Ошибка сохранения в БД", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newTask)

}
