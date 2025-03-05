package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func PatchHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Получен PATCH-запрос")

	vars := mux.Vars(r)
	id := vars["ID"]

	var existingTask Task
	if err := DB.First(&existingTask, id).Error; err != nil {
		http.Error(w, `{"error": "Задача не найдена"}`, http.StatusNotFound)
		return
	}

	var updateData struct {
		Task   *string `json:"task"`
		IsDone *bool   `json:"is_done"`
	}
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, `{"error": "Ошибка декодирования JSON"}`, http.StatusBadRequest)
		return
	}

	if updateData.Task != nil {
		existingTask.Task = *updateData.Task
	}

	if updateData.IsDone != nil {
		existingTask.IsDone = *updateData.IsDone
	}

	if err := DB.Save(&existingTask).Error; err != nil {
		http.Error(w, `{"error": "Ошибка обновления в БД"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existingTask)
}
