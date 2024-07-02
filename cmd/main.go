package main

import (
	"github.com/zsandibe/command-runner-service/internal/app"
	logger "github.com/zsandibe/command-runner-service/pkg"
)

func main() {
	if err := app.Start(); err != nil {
		logger.Error(err)
		return
	}
}
