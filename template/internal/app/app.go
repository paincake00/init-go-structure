package app

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	logger *slog.Logger
	cfg    *config.Config
}

func NewApp(logger *slog.Logger, cfg *config.Config) *App {
	app := &App{}

	app.logger = logger
	app.cfg = cfg

	return app
}

func (app *App) Run() error {
	errChan := make(chan error, 10)

	shutdown := make(chan error, 1) // канал для ошибок во время остановки
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		sig := <-quit

		app.logger.Info("Received signal", slog.String("signal", sig.String()))

		// Stop servers

		close(shutdown)
	}()

	select {
	case err := <-errChan:
		app.logger.Error("Fatal background error", sl.Err(err))
		return err
	case <-shutdown:
		app.logger.Info("Service stopped gracefully")
	}

	return nil
}
