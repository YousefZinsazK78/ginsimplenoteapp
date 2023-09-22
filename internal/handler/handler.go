package handler

import "github.com/yousefzinsazk78/simple_note_api/internal/database"

type Handler struct {
	noteStorer database.NoteStorer
}
