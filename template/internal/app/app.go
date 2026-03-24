package app

import (
	"go.uber.org/zap"
)

type App struct {
	logger *zap.SugaredLogger
	config Config
}

func NewApp(logger *zap.SugaredLogger, config Config) *App {
	app := &App{}

	app.logger = logger
	app.config = config

	return app
}

func (app *App) Run() error {
	return nil
}