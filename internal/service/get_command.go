package service

import (
	"context"

	"github.com/zsandibe/command-runner-service/internal/entity"
)

func (s *Service) GetCommandById(ctx context.Context, id int) (*entity.Command, error) {

	r, err := s.Command.GetCommandById(ctx, int64(id))
	if err != nil {
		return nil, err
	}

	err = s.ScriptCache.Set(r.Id, r.Script)
	if err != nil {
		return nil, err
	}

	return r, nil
}
