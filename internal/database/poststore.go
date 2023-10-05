package database

import (
	"context"

	"github.com/yousefzinsazk78/simple_note_api/internal/models"
)

type PostStorer interface {
	InsertPost(models.Post) error
	GetPosts() ([]models.Post, error)
	GetPostByTitle(string) (models.Post, error)
	DeletePost(int) error
	Update(models.Post) error
}

type postStore struct {
	database
}

func NewPostStore(db database) *postStore {
	return &postStore{
		database: db,
	}
}

func (d *noteStore) InsertPost(post models.Post) error {
	stmt, err := d.DB.Prepare("INSERT INTO notetbl(title,body) VALUES($1,$2)")

	if err != nil {
		return err
	}
	defer stmt.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err = stmt.ExecContext(ctx, note.Title, note.Body)
	if err != nil {
		return err
	}
	return nil
}
