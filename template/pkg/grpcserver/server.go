package grpcserver

import (
	"context"
	"errors"
	"log/slog"
	"net"

	"google.golang.org/grpc"
)

const (
	defaultGRPCServerAddress = ":44044"
)

type Server struct {
	address    string
	App        *grpc.Server
	serverOpts []grpc.ServerOption
	notify     chan error
	logger     *slog.Logger
}

func NewServer(logger *slog.Logger, opts ...Option) *Server {
	srv := &Server{
		address: defaultGRPCServerAddress,
		App:     nil,
		notify:  make(chan error, 1),
		logger:  logger,
	}

	for _, opt := range opts {
		opt(srv)
	}

	srv.App = grpc.NewServer(srv.serverOpts...)

	return srv
}

func (s *Server) Start() {
	go func() {
		lis, err := net.Listen("tcp", s.address)
		if err != nil {
			s.notify <- err

			close(s.notify)

			return
		}

		s.logger.Info("Starting gRPC Server on " + s.address)

		err = s.App.Serve(lis)
		// Проверка что буфер пуст
		if len(s.notify) == 0 {
			s.notify <- err
		}

		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	const op = "grpcserver.Shutdown"

	var shutdownErrors []error

	// Graceful stop gRPC Server
	s.App.GracefulStop()

	// Проверка на ошибки из notify (если был не interrupt)
	if err, ok := <-s.notify; ok {
		if err != nil && !errors.Is(err, context.Canceled) {
			s.logger.Error(op, err)
			shutdownErrors = append(shutdownErrors, err)
		}
	}

	return errors.Join(shutdownErrors...)
}