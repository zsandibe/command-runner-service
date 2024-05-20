package postgres

import "github.com/jmoiron/sqlx"

type CommandOutput struct {
	db *sqlx.DB
}

func NewCommandOutputRepo(db *sqlx.DB) *CommandOutput {
	return &CommandOutput{
		db: db,
	}
}
