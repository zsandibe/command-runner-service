package postgres

import (
	"github.com/jmoiron/sqlx"
)

type CommandRepo struct {
	db *sqlx.DB
}

func NewCommandRepo(db *sqlx.DB) *CommandRepo {
	return &CommandRepo{
		db: db,
	}
}
