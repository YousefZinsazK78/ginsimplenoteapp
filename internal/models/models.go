package models

import (
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
)

//id , post_id, author_id, body,created_at, updated_at, like

type Image struct {
	ID       int    `json:"id"`
	ImageUrl string `json:"image_url"`
	PostID   int    `json:"post_id"`
	UserID   int    `json:"user_id"`
}

type Comment struct {
	ID        int          `json:"id"`
	PostID    int          `json:"post_id"`
	AuthorID  int          `json:"author_id"`
	Body      string       `json:"body"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"udpated_at"`
}

type UpdateCommentParam struct {
	Body string `json:"body"`
}

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

type Register struct {
	Username string `json:"username,required"`
	Email    string `json:"email,required"`
	Password string `json:"password,required"`
}

type Login struct {
	Username string `json:"username,required"`
	Password string `json:"password,required"`
}

type User struct {
	ID        int          `json:"id"`
	RoleID    int          `json:"role_id"`
	Username  string       `json:"username"`
	Password  string       `json:"-"`
	Email     string       `json:"email"`
	CreatedAt time.Time    `json:"createdat"`
	UpdatedAt sql.NullTime `json:"updatedat"`
}

type UserRole struct {
	ID     int `json:"id"`
	RoleID int `json:"role_id"`
}

type Role struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// utils methods
func (user *User) HashPassword() error {
	passHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passHash)
	return nil
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
