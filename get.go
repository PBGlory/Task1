package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("Получен GET-запрос")
	var tasks []Task

	if err := DB.Find(&tasks).Error; err != nil {
		http.Error(writer, "Ошибка получения данных", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(writer).Encode(tasks); err != nil {
		http.Error(writer, "Ошибка кодирования ответа", http.StatusInternalServerError)
	}
}
