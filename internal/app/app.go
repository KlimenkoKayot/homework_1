package app

import (
	"github.com/klimenkokayot/homework_1/config"
	"go.uber.org/zap"
)

func NewApplication(logger *zap.Logger, cfg *config.Config) *Application {
	logger.Info("application init")
	return &Application{
		Logger: logger,
		Config: cfg,
	}
}

type Application struct {
	Logger *zap.Logger
	Config *config.Config
}

func (app *Application) Run() error {
	app.Logger.Info("run application")
	return nil
}
