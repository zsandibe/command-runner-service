package service

import (
	repo "github.com/zsandibe/command-runner-service/internal/repository"
)

type Command interface {
}

type CommandOutput interface {
}

type Service struct {
	Command
	CommandOutput
}

func NewService(repo *repo.Repository) *Service {
	return &Service{
		Command:       NewCommandService(repo.Command),
		CommandOutput: NewCommandOutputService(repo.CommandOutput),
	}
}
