package service

import repo "github.com/zsandibe/command-runner-service/internal/repository"

type CommandService struct {
	commandRepo repo.Command
}

func NewCommandService(commandRepo repo.Command) *CommandService {
	return &CommandService{commandRepo: commandRepo}
}
