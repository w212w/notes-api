package main

import (
	"log"
	"net/http"
	"notes-api/handlers"
	"notes-api/repository"

	"github.com/gorilla/mux"
)

func main() {
	repo := repository.NewNoteRepository()
	handler := handlers.NewNoteHandler(repo)

	router := mux.NewRouter()

	// Настройка маршрутов
	router.HandleFunc("/notes", handler.GetAllNotes).Methods("GET")
	router.HandleFunc("/notes/{id}", handler.GetNote).Methods("GET")
	router.HandleFunc("/notes", handler.CreateNote).Methods("POST")
	router.HandleFunc("/notes/{id}", handler.UpdateNote).Methods("PUT")
	router.HandleFunc("/notes/{id}", handler.DeleteNote).Methods("DELETE")
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	log.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
