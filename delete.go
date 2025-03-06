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

	var deleteData struct {
		Task   *string `json:"task"`
		IsDone *bool   `json:"is_done"`
	}

	if err := DB.Delete(&existingTask).Error; err != nil {
		http.Error(writer, `{"error": "Ошибка декодирования JSON"}`, http.StatusBadRequest)
	}

	if deleteData.Task != nil {
		existingTask.Task = *deleteData.Task
	}

	if deleteData.IsDone != nil {
		existingTask.IsDone = *deleteData.IsDone
	}

	if err := DB.Save(&existingTask).Error; err != nil {
		http.Error(writer, `{"error": "Ошибка обновления БД"}`, http.StatusInternalServerError)

	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(existingTask)
}
