package handler

import "github.com/yousefzinsazk78/simple_note_api/internal/database"

type handler struct {
	noteStorer database.NoteStorer
	postStorer database.PostStorer
}

func NewHandler(noteStorer database.NoteStorer, poststorer database.PostStorer) *handler {
	return &handler{
		noteStorer: noteStorer,
		postStorer: poststorer,
	}
}
