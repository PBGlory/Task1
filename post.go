package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func PostHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("Получен POST-запрос")

	var newTask Task

	if err := json.NewDecoder(request.Body).Decode(&newTask); err != nil {
		http.Error(writer, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	newTask.Task = strings.TrimSpace(newTask.Task)

	if len(newTask.Task) == 0 {
		http.Error(writer, "Поле 'task' обязательно и должно содержать символы", http.StatusBadRequest)
		return
	}

	if err := DB.Create(&newTask).Error; err != nil {
		http.Error(writer, "Ошибка сохранения в БД", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)

	json.NewEncoder(writer).Encode(newTask)

}
