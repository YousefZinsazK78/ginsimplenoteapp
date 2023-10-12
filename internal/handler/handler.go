package handler

import "github.com/yousefzinsazk78/simple_note_api/internal/database"

type handler struct {
	noteStorer    database.NoteStorer
	postStorer    database.PostStorer
	roleStorer    database.RoleStorer
	userStorer    database.UserStorer
	commentStorer database.CommentStorer
}

func NewHandler(noteStorer database.NoteStorer, poststorer database.PostStorer, roleStorer database.RoleStorer, userStorer database.UserStorer, commentStorer database.CommentStorer) *handler {
	return &handler{
		noteStorer:    noteStorer,
		postStorer:    poststorer,
		roleStorer:    roleStorer,
		userStorer:    userStorer,
		commentStorer: commentStorer,
	}
}
