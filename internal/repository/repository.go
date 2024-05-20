package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/zsandibe/command-runner-service/internal/repository/postgres"
)

type Command interface {
}

type CommandOutput interface {
}

type Repository struct {
	Command
	CommandOutput
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Command:       postgres.NewCommandRepo(db),
		CommandOutput: postgres.NewCommandOutputRepo(db),
	}
}
