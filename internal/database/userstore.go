package database

import (
	"context"

	"github.com/yousefzinsazk78/simple_note_api/internal/models"
)

type UserStorer interface {
	InsertUser(models.User) error
	GetUsers() ([]models.User, error)
	GetUserByID(int) (*models.User, error)
	UpdateUser(models.User) error
}

type userStore struct {
	database
}

func NewUserStore(db database) *userStore {
	return &userStore{
		database: db,
	}
}

func (u *userStore) InsertUser(user models.User) error {
	stmt, err := u.DB.Prepare("INSERT INTO usertbl(roleid, username, password, email) VALUES($1,$2,$3,$4)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err = stmt.ExecContext(ctx, user.RoleID, user.Username, user.Password, user.Email)
	if err != nil {
		return err
	}
	return nil
}
