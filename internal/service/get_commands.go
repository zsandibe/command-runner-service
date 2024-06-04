package service

import (
	"context"
	"strconv"

	"github.com/zsandibe/command-runner-service/internal/entity"
)

func (s *Service) GetCommands(ctx context.Context, ids []string) (*[]entity.Command, error) {
	arr := make([]int64, 0, len(ids))
	for _, id := range ids {
		i, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return nil, err
		}
		arr = append(arr, i)
	}

	commands, err := s.Command.GetCommands(ctx, arr)

	if err != nil {
		return nil, err
	}
	for _, command := range *commands {
		err = s.ScriptCache.Set(command.Id, command.Script)
		if err != nil {
			return nil, err
		}
	}

	return commands, nil
}
