package database

import "github.com/yousefzinsazk78/simple_note_api/internal/models"

type NoteStorer interface {
	InsertNote(models.Note) error
	GetNotes() ([]models.Note, error)
	GetNotesByTitle(string) (models.Note, error)
	DeleteNote(int) error
	UpdateNote(models.Note) error
}
