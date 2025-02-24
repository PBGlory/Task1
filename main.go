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
	fmt.Fprintln(w, "Hello,", task)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received POST request")

	var requestBody struct {
		Message string `json:"message"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if requestBody.Message == "" {
		http.Error(w, "Message is required", http.StatusBadRequest)
		return
	}
	task = requestBody.Message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message": "Message updated to: %s"}`, task)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", GetHandler).Methods("GET")
	router.HandleFunc("/api/hello", PostHandler).Methods("POST")

	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
