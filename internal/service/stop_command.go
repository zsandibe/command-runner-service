package service

import (
	"fmt"
	"os/exec"
)

func (s *Service) StopCommandById(id int) error {
	cmd, err := s.ExecCmdCache.Get(int64(id))
	if err != nil {
		return err
	}
	c := cmd.(*exec.Cmd)
	if err = c.Process.Kill(); err != nil {
		return fmt.Errorf("failed to kill process: %e", err)
	}
	return nil
}
