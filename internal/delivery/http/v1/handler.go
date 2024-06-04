package v1

import (
	"context"

	"github.com/zsandibe/command-runner-service/internal/domain"
	"github.com/zsandibe/command-runner-service/internal/entity"
)

//go:generate mockgen -destination=mocks/handler.go -package=mock -source=handler.go
//go:generate touch mocks/.coverignore

type Service interface {
	CreateCommand(ctx context.Context, inp *domain.CreateCommandInput) (*entity.Command, error)
	GetCommandById(ctx context.Context, id int) (*entity.Command, error)
	GetAllCommands(ctx context.Context) (*[]entity.Command, error)
	GetCommands(ctx context.Context, ids []string) (*[]entity.Command, error)
	DeleteCommandById(ctx context.Context, id int) error
	StopCommandById(id int) error
}

type Handler struct {
	Service
}

func NewHandler(service Service) *Handler {
	return &Handler{Service: service}
}
