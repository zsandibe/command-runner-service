package service

import (
	"context"

	"github.com/zsandibe/command-runner-service/internal/domain"
	"github.com/zsandibe/command-runner-service/internal/entity"
)

func (s *Service) CreateCommand(ctx context.Context, inp *domain.CreateCommandInput) (*entity.Command, error) {
	return s.Command.CreateCommand(ctx, *inp)
}
