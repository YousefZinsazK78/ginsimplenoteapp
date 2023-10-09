package handler

import "github.com/yousefzinsazk78/simple_note_api/internal/database"

type handler struct {
	noteStorer database.NoteStorer
	postStorer database.PostStorer
	roleStorer database.RoleStorer
	userStorer database.UserStorer
}

func NewHandler(noteStorer database.NoteStorer, poststorer database.PostStorer, roleStorer database.RoleStorer, userStorer database.UserStorer) *handler {
	return &handler{
		noteStorer: noteStorer,
		postStorer: poststorer,
		roleStorer: roleStorer,
		userStorer: userStorer,
	}
}
