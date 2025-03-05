package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Получен GET-запрос")
	var tasks []Task

	if err := DB.Find(&tasks).Error; err != nil {
		http.Error(w, "Ошибка получения данных", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, "Ошибка кодирования ответа", http.StatusInternalServerError)
	}
}
