package service

import (
	"context"

	"github.com/zsandibe/command-runner-service/internal/entity"
)

func (s *Service) GetAllCommands(ctx context.Context) (*[]entity.Command, error) {
	commands, err := s.Command.GetAllCommands(ctx)
	if err != nil {
		return nil, err
	}

	return commands, nil
}
