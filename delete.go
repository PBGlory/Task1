package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func DeleteHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("Получение DELETE-запрос")

	vars := mux.Vars(request)
	id := vars["id"]

	var existingTask Task

	if err := DB.First(&existingTask, id).Error; err != nil {
		http.Error(writer, `{"error": "Задача не найдена"}`, http.StatusNotFound)
		return
	}

	DB.Delete(&existingTask)

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(existingTask)
}
