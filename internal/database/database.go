package database

import "database/sql"

type database struct {
	*sql.DB
}

func NewDatabase(db *sql.DB) *database {
	return &database{
		DB: db,
	}
}
