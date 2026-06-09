package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world\n"))
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
