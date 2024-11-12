package repository

import "notes-api/models"

type NoteRepository struct {
	notes  []models.Note
	nextID int
}

// NewNoteRepository создает новый репозиторий с начальными данными
func NewNoteRepository() *NoteRepository {
	return &NoteRepository{
		notes:  []models.Note{},
		nextID: 1,
	}
}

// GetAll возвращает все заметки
func (r *NoteRepository) GetAll() []models.Note {
	return r.notes
}

// GetByID возвращает заметку по ID
func (r *NoteRepository) GetByID(id int) *models.Note {
	for _, note := range r.notes {
		if note.ID == id {
			return &note
		}
	}
	return nil
}

// Create добавляет новую заметку
func (r *NoteRepository) Create(note models.Note) models.Note {
	note.ID = r.nextID
	r.nextID++
	r.notes = append(r.notes, note)
	return note
}

// Update обновляет существующую заметку
func (r *NoteRepository) Update(id int, updatedNote models.Note) *models.Note {
	for i, note := range r.notes {
		if note.ID == id {
			r.notes[i].Title = updatedNote.Title
			r.notes[i].Content = updatedNote.Content
			return &r.notes[i]
		}
	}
	return nil
}

// Delete удаляет заметку по ID
func (r *NoteRepository) Delete(id int) bool {
	for i, note := range r.notes {
		if note.ID == id {
			r.notes = append(r.notes[:i], r.notes[i+1:]...)
			return true
		}
	}
	return false
}
