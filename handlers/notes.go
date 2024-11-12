package handlers

import (
	"encoding/json"
	"net/http"
	"notes-api/models"
	"notes-api/repository"
	"strconv"

	"github.com/gorilla/mux"
)

type NoteHandler struct {
	Repo *repository.NoteRepository
}

// NewNoteHandler создает новый обработчик заметок
func NewNoteHandler(repo *repository.NoteRepository) *NoteHandler {
	return &NoteHandler{Repo: repo}
}

// GetAllNotes обрабатывает запрос на получение всех заметок
func (h *NoteHandler) GetAllNotes(w http.ResponseWriter, r *http.Request) {
	notes := h.Repo.GetAll()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

// GetNote обрабатывает запрос на получение заметки по ID
func (h *NoteHandler) GetNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	note := h.Repo.GetByID(id)
	if note == nil {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

// CreateNote обрабатывает запрос на создание новой заметки
func (h *NoteHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var note models.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	createdNote := h.Repo.Create(note)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdNote)
}

// UpdateNote обрабатывает запрос на обновление заметки
func (h *NoteHandler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	var updatedNote models.Note
	if err := json.NewDecoder(r.Body).Decode(&updatedNote); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	note := h.Repo.Update(id, updatedNote)
	if note == nil {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

// DeleteNote обрабатывает запрос на удаление заметки
func (h *NoteHandler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	if !h.Repo.Delete(id) {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
