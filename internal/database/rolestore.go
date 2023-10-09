package database

import (
	"context"

	"github.com/yousefzinsazk78/simple_note_api/internal/models"
)

type RoleStorer interface {
	InsertRole(models.Role) error
	GetRoles() ([]models.Role, error)
	GetRoleByID(int) (*models.Role, error)
	UpdateRole(models.Role) error
}

type roleStore struct {
	database
}

func NewRoleStore(db database) *roleStore {
	return &roleStore{
		database: db,
	}
}

func (u *roleStore) InsertRole(role models.Role) error {
	stmt, err := u.DB.Prepare("INSERT INTO roletbl(name, description) VALUES($1,$2)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err = stmt.ExecContext(ctx, role.Name, role.Description)
	if err != nil {
		return err
	}
	return nil
}

func (d *roleStore) GetRoles() ([]models.Role, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	rows, err := d.DB.QueryContext(ctx, "SELECT * FROM roletbl;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roleModels []models.Role
	for rows.Next() {
		var roleM models.Role
		if err := rows.Scan(&roleM.ID, &roleM.Name, &roleM.Description); err != nil {
			return nil, err
		}
		roleModels = append(roleModels, roleM)
	}

	return roleModels, nil
}

func (d *roleStore) GetRoleByID(id int) (*models.Role, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	row := d.DB.QueryRowContext(ctx, "SELECT * FROM roletbl WHERE id=$1", id)
	var roleM models.Role
	if err := row.Scan(&roleM.ID, &roleM.Name, &roleM.Description); err != nil {
		return nil, err
	}
	return &roleM, nil
}

func (d *roleStore) UpdateRole(role models.Role) error {
	rootCtx := context.Background()
	ctx, cancel := context.WithCancel(rootCtx)
	defer cancel()
	stmt, err := d.DB.PrepareContext(ctx, "UPDATE roletbl SET name=$1 WHERE ID=$2;")
	if err != nil {
		return err
	}
	defer stmt.Close()
	secondCtx, secondCancel := context.WithCancel(rootCtx)
	defer secondCancel()
	_, err = stmt.ExecContext(secondCtx, role.Name, role.ID)
	if err != nil {
		return err
	}
	return nil
}
