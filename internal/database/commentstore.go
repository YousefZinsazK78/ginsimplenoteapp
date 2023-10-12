package database

import (
	"context"

	"github.com/yousefzinsazk78/simple_note_api/internal/models"
)

type CommentStorer interface {
	Insert(models.Comment) error
	GetComments() ([]models.Comment, error)
	GetCommentByID(int) (models.Comment, error)
	GetCommentByBody(string) (models.Comment, error)
	Update(int, models.UpdateCommentParam) error
	Delete(int) error
}

type commentStore struct {
	database
}

func NewCommentStore(db database) *commentStore {
	return &commentStore{
		database: db,
	}
}

func (c *commentStore) Insert(comment models.Comment) error {
	stmt, err := c.DB.Prepare("INSERT INTO commenttbl(PostID,AuthorID,body) VALUES($1,$2,$3)")

	if err != nil {
		return err
	}
	defer stmt.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err = stmt.ExecContext(ctx, comment.PostID, comment.AuthorID, comment.Body)
	if err != nil {
		return err
	}
	return nil
}
