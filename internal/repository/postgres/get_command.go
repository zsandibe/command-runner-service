package postgres

import (
	"context"

	"github.com/zsandibe/command-runner-service/internal/entity"
)

func (r *CommandRepo) GetCommandById(ctx context.Context, id int64) (*entity.Command, error) {
	var command entity.Command

	query := `
		SELECT command.id, command.script, command.description FROM command WHERE id = $1
	`

	if err := r.db.QueryRowContext(ctx, query, id).Scan(&command.Id, &command.Script, &command.Description); err != nil {
		return &command, err
	}
	return &command, nil
}
