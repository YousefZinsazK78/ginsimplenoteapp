package models

import (
	"database/sql"
	"time"
)

type Note struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type Post struct {
	ID        int          `json:"id"`
	Title     string       `json:"title"`
	Subtitle  string       `json:"subtitle"`
	Body      string       `json:"body"`
	AuthorID  int          `json:"authorid"`
	View      int          `json:"view"`
	CreatedAt time.Time    `json:"createdat"`
	UpdatedAt sql.NullTime `json:"updatedat"`
}
