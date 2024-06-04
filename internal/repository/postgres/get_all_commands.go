package postgres

import (
	"context"

	"github.com/zsandibe/command-runner-service/internal/entity"
)

func (r *CommandRepo) GetAllCommands(ctx context.Context) (*[]entity.Command, error) {
	var cs []entity.Command
	query := "SELECT id, script, description FROM commands"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		c := entity.Command{}
		err = rows.Scan(&c.Id, &c.Script, &c.Description)
		if err != nil {
			return nil, err
		}
		cs = append(cs, c)
	}
	return &cs, nil
}
