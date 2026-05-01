package main

import (
	"log"
	"log/slog"
)

// Example run: CONFIG_PATH=".env" go run cmd/mainapp/main.go

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config", sl.Err(err))
	}

	logger := logs.SetupLogger(cfg.Env)

	logger.Info("Starting IAM Service with config:", slog.Any("config", cfg))

	application := app.NewApp(logger, cfg)

	if errApp := application.Run(); errApp != nil {
		logger.Error("Failed to start application", sl.Err(errApp))
	}
}
