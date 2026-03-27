package main

import (
	"log"
	"log/slog"
)

func main() {
	cfg, err := config.MustLoad()
	if err != nil {
		log.Fatal("failed to load config", sl.Err(err))
	}

	logger := logs.SetupLogger(cfg.Env)
}