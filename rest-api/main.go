package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HealthResponse struct {
	Status string `json:"status"`
}

type Book struct {
	Id       string   `json:"id"`
	Title    string   `json:"title"`
	Author   string   `json:"author"`
	Year     int      `json:"year"`
	Category []string `json:"category"`
}

type BookList []Book

var books BookList

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world\n"))
	})

	router.HandleFunc("GET /books", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("All Books\n"))
	})

	router.HandleFunc("POST /books", func(w http.ResponseWriter, r *http.Request) {
		// Stopping here. I need to implement this endpoint

		w.Write([]byte("Create book\n"))
	})

	router.HandleFunc("GET /books/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		w.Write([]byte(fmt.Sprintf("id: %s\n", id)))
	})

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		healthyStatus := HealthResponse{Status: "ok"}
		status, err := json.Marshal(healthyStatus)
		if err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(status)
	})

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Serving the server failed: ", err)
	}
}
