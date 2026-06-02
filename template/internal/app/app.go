package app

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

// useCases (или service), описывающие внутреннюю бизнес логику
type useCases struct {
}

// servers содержит серверы/воркеры, которые обращаются к внутренним useCases
type servers struct {
}

func Run(cfg *config.Config) {
	// Внешние зависимости (логгеры, конфиги, подключения)
	logger := logs.SetupLogger(cfg.Env)

	logger.Info("Starting service with config:", slog.Any("config", cfg))

	uc := initUseCases(logger, cfg)

	s := initServers(logger, cfg, uc)

	s.startServers()

	s.waitServersForShutdown(logger)
}

func initUseCases(l *slog.Logger, cfg *config.Config) *useCases {
	return &useCases{}
}

func initServers(l *slog.Logger, cfg *config.Config, uc *useCases) *servers {
	// При ошибках (при обработке) можно вызывать os.Exit(1)

	return &servers{}
}

func (s *servers) startServers() {

}

func (s *servers) waitServersForShutdown(l *slog.Logger) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-interrupt:
		l.Info("Got interrupt signal for graceful shutdown: ", slog.String("signal", sig.String()))
	}

	s.shutdownServers(l)
}

func (s *servers) shutdownServers(l *slog.Logger) {

}
