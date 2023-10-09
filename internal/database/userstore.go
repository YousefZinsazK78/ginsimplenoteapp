package database

import (
	"context"

	"github.com/yousefzinsazk78/simple_note_api/internal/models"
)

type UserStorer interface {
	InsertUser(models.User) error
	GetUsers() ([]models.User, error)
	GetUserByID(int) (*models.User, error)
	UpdateUser(models.UserRole) error
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

func (d *userStore) GetUsers() ([]models.User, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	rows, err := d.DB.QueryContext(ctx, "SELECT * FROM usertbl;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userModels []models.User
	for rows.Next() {
		var userM models.User
		if err := rows.Scan(&userM.ID, &userM.RoleID, &userM.Username, &userM.Password, &userM.Email, &userM.CreatedAt, &userM.UpdatedAt); err != nil {
			return nil, err
		}
		userModels = append(userModels, userM)
	}

	return userModels, nil
}

func (d *userStore) GetUserByID(id int) (*models.User, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	row := d.DB.QueryRowContext(ctx, "SELECT * FROM usertbl WHERE id=$1", id)
	var userM models.User
	if err := row.Scan(&userM.ID, &userM.RoleID, &userM.Username, &userM.Password, &userM.Email, &userM.CreatedAt, &userM.UpdatedAt); err != nil {
		return nil, err
	}
	return &userM, nil
}

func (d *userStore) UpdateUser(userRole models.UserRole) error {
	rootCtx := context.Background()
	ctx, cancel := context.WithCancel(rootCtx)
	defer cancel()
	stmt, err := d.DB.PrepareContext(ctx, "UPDATE usertbl SET roleid=$1 WHERE ID=$2;")
	if err != nil {
		return err
	}
	defer stmt.Close()
	secondCtx, secondCancel := context.WithCancel(rootCtx)
	defer secondCancel()
	_, err = stmt.ExecContext(secondCtx, userRole.RoleID, userRole.ID)
	if err != nil {
		return err
	}
	return nil
}
