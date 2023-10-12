package database

import (
	"context"
	"time"

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

// Insert func Insert Comment to CommentTbl
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

func (c *commentStore) GetComments() ([]models.Comment, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	rows, err := c.DB.QueryContext(ctx, "SELECT * FROM commenttbl;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var commentModels []models.Comment
	for rows.Next() {
		var commentM models.Comment
		if err := rows.Scan(&commentM.ID, &commentM.PostID, &commentM.AuthorID, &commentM.Body, &commentM.CreatedAt, &commentM.UpdatedAt); err != nil {
			return nil, err
		}
		commentModels = append(commentModels, commentM)
	}

	return commentModels, nil
}

func (c *commentStore) GetCommentByID(id int) (models.Comment, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	row := c.DB.QueryRowContext(ctx, "SELECT * FROM commenttbl WHERE id=$1", id)
	var commentM models.Comment
	if err := row.Scan(&commentM.ID, &commentM.PostID, &commentM.AuthorID, &commentM.Body, &commentM.CreatedAt, &commentM.UpdatedAt); err != nil {
		return models.Comment{}, err
	}
	return commentM, nil
}

func (c *commentStore) GetCommentByBody(body string) (models.Comment, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	row := c.DB.QueryRowContext(ctx, "SELECT * FROM commenttbl WHERE body LIKE $1", body)
	var commentM models.Comment
	if err := row.Scan(&commentM.ID, &commentM.PostID, &commentM.AuthorID, &commentM.Body, &commentM.CreatedAt, &commentM.UpdatedAt); err != nil {
		return models.Comment{}, err
	}
	return commentM, nil
}

func (c *commentStore) Update(id int, updateParam models.UpdateCommentParam) error {
	rootCtx := context.Background()
	ctx, cancel := context.WithCancel(rootCtx)
	defer cancel()
	stmt, err := c.DB.PrepareContext(ctx, "UPDATE commenttbl SET body=$1,UpdatedAt=$2 WHERE ID=$3;")
	if err != nil {
		return err
	}
	defer stmt.Close()
	secondCtx, secondCancel := context.WithCancel(rootCtx)
	defer secondCancel()
	_, err = stmt.ExecContext(secondCtx, updateParam.Body, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}

func (c *commentStore) Delete(id int) error {
	rootCtx := context.Background()
	ctx, cancel := context.WithCancel(rootCtx)
	defer cancel()
	stmt, err := c.DB.PrepareContext(ctx, "DELETE FROM commenttbl WHERE ID=$1;")
	if err != nil {
		return err
	}
	defer stmt.Close()
	secondCtx, secondCancel := context.WithCancel(rootCtx)
	defer secondCancel()
	_, err = stmt.ExecContext(secondCtx, id)
	if err != nil {
		return err
	}
	return nil
}
