package service

import (
	"context"

	"github.com/zsandibe/command-runner-service/internal/domain"
	"github.com/zsandibe/command-runner-service/internal/entity"
)

//go:generate mockgen -destination=mocks/service.go -package=mock -source=service.go
//go:generate touch mocks/.coverignore

type Command interface {
	CreateCommand(ctx context.Context, inp domain.CreateCommandInput) (*entity.Command, error)
	GetCommandById(ctx context.Context, id int64) (*entity.Command, error)
	GetCommands(ctx context.Context, ids []int64) (*[]entity.Command, error)
	GetAllCommands(ctx context.Context) (*[]entity.Command, error)
	DeleteCommandById(ctx context.Context, id int64) error
	CreateCommandOutput(ctx context.Context, id int64, output string) error
}

type Cache interface {
	Set(key int64, value any) error
	Get(key int64) (any, error)
	GetAllKeys() ([]int64, error)
	Delete(key int64) error
	GetLen() (int, error)
}

type Service struct {
	Command
	ScriptCache  Cache
	ExecCmdCache Cache
	stopSignal   chan struct{}
}

func NewService(repo Command, scriptCache, execCmdCache Cache) *Service {
	stopSignal := make(chan struct{})

	s := &Service{
		repo,
		scriptCache,
		execCmdCache,
		stopSignal,
	}

	go s.Runner()
	return s
}
