package service

import repo "github.com/zsandibe/command-runner-service/internal/repository"

type CommandOutputService struct {
	outputRepo repo.CommandOutput
}

func NewCommandOutputService(outputRepo repo.CommandOutput) *CommandOutputService {
	return &CommandOutputService{outputRepo: outputRepo}
}
