package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"

	"github.com/google/uuid"
)

//
// Types
//

type Book struct {
	ID       string   `json:"id"`
	Title    string   `json:"title"`
	Author   string   `json:"author"`
	Year     int      `json:"year"`
	Category []string `json:"category"`
}

func (b Book) String() string {
	return fmt.Sprintf("ID: %s\nTitle: %s\nAuthor: %s", b.ID, b.Title, b.Author)
}

type BookList []Book

var books BookList

//
// Responses
//

type HealthResponse struct {
	Status string `json:"status"`
}

type BookResponse struct {
	Status string `json:"status"`
	Book   *Book  `json:"book,omitempty"`
}

type BooksResponse struct {
	Status string   `json:"status"`
	Books  BookList `json:"books"`
}

//
// Main
//

func writeResponse(w http.ResponseWriter, status int, data any) {
	response, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		healthyStatus := HealthResponse{Status: "ok"}

		writeResponse(w, http.StatusOK, healthyStatus)
	})

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world\n"))
	})

	router.HandleFunc("POST /books", func(w http.ResponseWriter, r *http.Request) {
		book := Book{}
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			http.Error(w, "failed to decode body", http.StatusBadRequest)
			return
		}
		book.ID = uuid.New().String()

		books = append(books, book)

		responseRaw := BookResponse{
			Status: "ok",
			Book:   &book,
		}
		writeResponse(w, http.StatusCreated, responseRaw)
	})

	router.HandleFunc("GET /books", func(w http.ResponseWriter, r *http.Request) {
		responseRaw := BooksResponse{
			Status: "ok",
			Books:  books,
		}
		writeResponse(w, http.StatusOK, responseRaw)
	})

	router.HandleFunc("GET /books/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		for _, book := range books {
			if book.ID == id {

				responseRaw := BookResponse{
					Status: "ok",
					Book:   &book,
				}
				writeResponse(w, http.StatusOK, responseRaw)
				return
			}
		}
		responseRaw := BookResponse{
			Status: "book not found",
		}
		writeResponse(w, http.StatusNotFound, responseRaw)
	})

	router.HandleFunc("PUT /books/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		updatedBook := Book{}
		err := json.NewDecoder(r.Body).Decode(&updatedBook)
		if err != nil {
			http.Error(w, "failed to decode body", http.StatusBadRequest)
			return
		}
		updatedBook.ID = id

		for index, book := range books {
			if book.ID == id {
				books[index] = updatedBook
				responseRaw := BookResponse{
					Status: "ok",
					Book:   &updatedBook,
				}
				writeResponse(w, http.StatusOK, responseRaw)
				return
			}
		}
		writeResponse(w, http.StatusNotFound, BookResponse{Status: "book not found"})
	})

	router.HandleFunc("DELETE /books/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		books = slices.DeleteFunc(books, func(book Book) bool { return book.ID == id })
		writeResponse(w, http.StatusOK, BookResponse{Status: "ok"})
	})

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Serving the server failed: ", err)
	}
}
