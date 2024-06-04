package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/zsandibe/command-runner-service/internal/domain"
	"github.com/zsandibe/command-runner-service/internal/entity"
	"github.com/zsandibe/command-runner-service/internal/repository/postgres"
)

type Command interface {
	CreateCommand(ctx context.Context, inp domain.CreateCommandInput) (*entity.Command, error)
	GetCommandById(ctx context.Context, id int64) (*entity.Command, error)
	GetCommands(ctx context.Context, ids []int64) (*[]entity.Command, error)
	GetAllCommands(ctx context.Context) (*[]entity.Command, error)
	DeleteCommandById(ctx context.Context, id int64) error
	CreateCommandOutput(ctx context.Context, id int64, output string) error
}

type Repository struct {
	Command
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Command: postgres.NewCommandRepo(db),
	}
}
