package handler

import "github.com/yousefzinsazk78/simple_note_api/internal/database"

type handler struct {
	noteStorer database.NoteStorer
}

func NewHandler(noteStorer database.NoteStorer) *handler {
	return &handler{
		noteStorer: noteStorer,
	}
}
