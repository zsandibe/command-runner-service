package postgres

import "context"

func (r *CommandRepo) DeleteCommandById(ctx context.Context, id int64) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	deleteOutputsCommandQuery := "DELETE FROM commands_output WHERE id_command = $1"
	_, err = tx.ExecContext(ctx, deleteOutputsCommandQuery, id)

	if err != nil {
		_ = tx.Rollback()
		return err
	}

	deleteCommandQuery := "DELETE FROM commands WHERE id = $1"
	_, err = tx.ExecContext(ctx, deleteCommandQuery, id)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
