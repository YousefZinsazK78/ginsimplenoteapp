package models

import (
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
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
