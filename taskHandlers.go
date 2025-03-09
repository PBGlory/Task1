package handlers

import (
	"Task1/internal/taskService"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Handler struct {
	Service *taskService.TaskService
}

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{Service: service}
}

func (handler *Handler) GetTaskHandler(writer http.ResponseWriter, request *http.Request) {
	tasks, err := handler.Service.GetAllTasks()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(tasks)
}

func (handler *Handler) PostTaskHandler(writer http.ResponseWriter, request *http.Request) {
	var task taskService.Task
	err := json.NewDecoder(request.Body).Decode(&task)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	createdTask, err := handler.Service.CreateTask(task)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(createdTask)
}

func (handler *Handler) UpdateTaskHandler(writer http.ResponseWriter, request *http.Request) {
	var task taskService.Task

	vars := mux.Vars(request)
	id := vars["id"]

	taskID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(writer, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(request.Body).Decode(&task)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	task.ID = uint(taskID)

	updatedTask, err := handler.Service.UpdateTaskByID(uint(taskID), task)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(updatedTask)
}

func (handler *Handler) DeleteTaskHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	taskID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(writer, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = handler.Service.DeleteTaskById(uint(taskID))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}
