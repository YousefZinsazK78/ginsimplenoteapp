package database

import (
	"context"
	"log"
	"time"

	"github.com/yousefzinsazk78/simple_note_api/internal/models"
)

type NoteStorer interface {
	InsertNote(models.Note) error
	GetNotes() ([]models.Note, error)
	GetNotesByTitle(string) (*models.Note, error)
	DeleteNote(int) error
	UpdateNote(models.Note) error
}

type noteStore struct {
	database
}

func NewNoteStore(db database) *noteStore {
	return &noteStore{
		database: db,
	}
}

func (d *noteStore) InsertNote(note models.Note) error {
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

func (d *noteStore) GetNotes() ([]models.Note, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	rows, err := d.DB.QueryContext(ctx, "SELECT * FROM notetbl;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var noteModels []models.Note
	for rows.Next() {
		var noteM models.Note
		if err := rows.Scan(&noteM.ID, &noteM.Title, &noteM.Body, &noteM.CreatedAt, &noteM.UpdatedAt); err != nil {
			return nil, err
		}
		noteModels = append(noteModels, noteM)
	}

	return noteModels, nil
}

func (d *noteStore) GetNotesByTitle(title string) (*models.Note, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	row := d.DB.QueryRowContext(ctx, "SELECT * FROM notetbl WHERE TITLE=$1", title)
	var noteM models.Note
	if err := row.Scan(&noteM.ID, &noteM.Title, &noteM.Body, &noteM.CreatedAt, &noteM.UpdatedAt); err != nil {
		return nil, err
	}
	return &noteM, nil
}

func (d *noteStore) DeleteNote(id int) error {
	rootCtx := context.Background()
	ctx, cancel := context.WithCancel(rootCtx)
	defer cancel()
	stmt, err := d.DB.PrepareContext(ctx, "DELETE FROM notetbl WHERE ID=$1;")
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

func (d *noteStore) UpdateNote(note models.Note) error {
	log.Println(note)
	rootCtx := context.Background()
	ctx, cancel := context.WithCancel(rootCtx)
	defer cancel()
	stmt, err := d.DB.PrepareContext(ctx, "UPDATE notetbl SET TITLE=$1, BODY=$2, UPDATED_AT=$3 WHERE ID=$4;")
	if err != nil {
		return err
	}
	defer stmt.Close()
	secondCtx, secondCancel := context.WithCancel(rootCtx)
	defer secondCancel()
	_, err = stmt.ExecContext(secondCtx, note.Title, note.Body, time.Now(), note.ID)
	if err != nil {
		return err
	}
	return nil
}
