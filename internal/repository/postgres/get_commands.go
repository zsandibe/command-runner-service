package postgres

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/zsandibe/command-runner-service/internal/entity"
)

func (r *CommandRepo) GetCommands(ctx context.Context, ids []int64) (*[]entity.Command, error) {
	var cs []entity.Command
	query, args, err := sqlx.In("SELECT id, script, description FROM commands WHERE id IN(?)", ids)
	if err != nil {
		return nil, err
	}

	query = sqlx.Rebind(sqlx.DOLLAR, query)
	rows, err := r.db.Query(query, args...)
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

	if len(cs) == 0 {
		return nil, sql.ErrNoRows
	}
	return &cs, nil
}
