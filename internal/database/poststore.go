package database

import (
	"context"
	"log"
	"time"

	"github.com/yousefzinsazk78/simple_note_api/internal/models"
)

type PostStorer interface {
	InsertPost(models.Post) error
	InsertImage(models.Image) error
	GetPosts() ([]models.Post, error)
	GetPostByTitle(string) (*models.Post, error)
	DeletePost(int) error
	UpdatePost(models.Post) error
}

type postStore struct {
	database
}

func NewPostStore(db database) *postStore {
	return &postStore{
		database: db,
	}
}

func (d *postStore) InsertPost(post models.Post) error {
	stmt, err := d.DB.Prepare("INSERT INTO posttbl(title,subtitle,body,authorID,imgurl) VALUES($1,$2,$3,$4,$5)")

	if err != nil {
		return err
	}
	defer stmt.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err = stmt.ExecContext(ctx, post.Title, post.Subtitle, post.Body, post.AuthorID, post.ImgUrl)
	if err != nil {
		return err
	}
	return nil
}

func (d *postStore) GetPosts() ([]models.Post, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	rows, err := d.DB.QueryContext(ctx, "SELECT * FROM posttbl;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var postModels []models.Post
	for rows.Next() {
		var postM models.Post
		if err := rows.Scan(&postM.ID, &postM.Title, &postM.Subtitle, &postM.Body, &postM.AuthorID, &postM.ImgUrl, &postM.CreatedAt, &postM.UpdatedAt); err != nil {
			return nil, err
		}
		log.Println(postM)
		postModels = append(postModels, postM)
	}

	return postModels, nil
}

func (d *postStore) GetPostByTitle(title string) (*models.Post, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	row := d.DB.QueryRowContext(ctx, "SELECT * FROM posttbl WHERE TITLE=$1", title)
	var postM models.Post
	if err := row.Scan(&postM.ID, &postM.Title, &postM.Subtitle, &postM.Body, &postM.AuthorID, &postM.ImgUrl, &postM.CreatedAt, &postM.UpdatedAt); err != nil {
		return nil, err
	}
	return &postM, nil
}

func (d *postStore) DeletePost(id int) error {
	rootCtx := context.Background()
	ctx, cancel := context.WithCancel(rootCtx)
	defer cancel()
	stmt, err := d.DB.PrepareContext(ctx, "DELETE FROM posttbl WHERE ID=$1;")
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

func (d *postStore) UpdatePost(post models.Post) error {
	rootCtx := context.Background()
	ctx, cancel := context.WithCancel(rootCtx)
	defer cancel()
	stmt, err := d.DB.PrepareContext(ctx, "UPDATE posttbl SET TITLE=$1,SUBTITLE=$2, BODY=$3, UPDATED_AT=$4 WHERE ID=$5;")
	if err != nil {
		return err
	}
	defer stmt.Close()
	secondCtx, secondCancel := context.WithCancel(rootCtx)
	defer secondCancel()
	_, err = stmt.ExecContext(secondCtx, post.Title, post.Subtitle, post.Body, time.Now(), post.ID)
	if err != nil {
		return err
	}
	return nil
}

func (d *postStore) InsertImage(image models.Image) error {
	stmt, err := d.DB.Prepare("INSERT INTO imagetbl(img_url, post_id, user_id) VALUES($1,$2,$3)")

	if err != nil {
		return err
	}
	defer stmt.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err = stmt.ExecContext(ctx, image.ImageUrl, image.PostID, image.UserID)
	if err != nil {
		return err
	}
	return nil
}
