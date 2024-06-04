package postgres

import (
	"context"

	"github.com/zsandibe/command-runner-service/internal/domain"
	"github.com/zsandibe/command-runner-service/internal/entity"
	errors "github.com/zsandibe/command-runner-service/internal/repository/repo_errors"
	logger "github.com/zsandibe/command-runner-service/pkg"
)

func (r *CommandRepo) CreateCommand(ctx context.Context, inp domain.CreateCommandInput) (*entity.Command, error) {
	var id int64
	var command *entity.Command
	query := `
		INSERT INTO command (script,description) VALUES ($1,$2) RETURNING id
	`

	err := r.db.QueryRowContext(ctx, query, inp.Script, inp.Description).Scan(&id)
	if err != nil {
		logger.Errorf("Error in creating command: %v", err)
		return nil, errors.ErrCreatingCommand
	}
	command = &entity.Command{
		Id:          id,
		Script:      inp.Script,
		Description: inp.Description,
	}
	return command, nil
}
