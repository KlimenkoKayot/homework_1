package main

import (
	"log"

	"github.com/klimenkokayot/homework_1/config"
	"github.com/klimenkokayot/homework_1/internal/app"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
		return
	}

	cfg, err := config.Load()
	if err != nil {
		logger.Error("cant load config", zap.Error(err))
		return
	}

	app := app.NewApplication(logger, cfg)
	if err := app.Run(); err != nil {
		logger.Error("error in runtime app", zap.Error(err))
	}
}
