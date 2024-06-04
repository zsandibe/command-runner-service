package service

import (
	"context"
)

func (s *Service) DeleteCommandById(ctx context.Context, id int) error {
	err := s.Command.DeleteCommandById(ctx, int64(id))
	if err != nil {
		return err
	}

	return nil
}
