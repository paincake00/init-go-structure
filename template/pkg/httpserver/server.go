package httpserver

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"
)

const (
	defaultAddress         = ":8080"
	defaultReadTimeout     = time.Second * 10
	defaultWriteTimeout    = time.Second * 10
	defaultShutdownTimeout = time.Second * 3
)

type Server struct {
	App             *http.Server
	address         string
	notify          chan error
	readTimeout     time.Duration
	writeTimeout    time.Duration
	shutdownTimeout time.Duration
	logger          *slog.Logger
}

func New(logger *slog.Logger, router http.Handler, opts ...Option) *Server {
	srv := &Server{
		App:             nil,
		address:         defaultAddress,
		notify:          make(chan error, 1),
		readTimeout:     defaultReadTimeout,
		writeTimeout:    defaultWriteTimeout,
		shutdownTimeout: defaultShutdownTimeout,
		logger:          logger,
	}

	for _, opt := range opts {
		opt(srv)
	}

	srv.App = &http.Server{
		Addr:         srv.address,
		Handler:      router,
		ReadTimeout:  srv.readTimeout,
		WriteTimeout: srv.writeTimeout,
	}

	return srv
}

func (s *Server) Start() {
	go func() {
		s.logger.Info("Starting http server on " + s.address)
		err := s.App.ListenAndServe()

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
	const op = "httpserver.Shutdown"

	var shutdownErrors []error

	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	// Завершение сервера и отлов ошибки при завершении
	err := s.App.Shutdown(ctx)
	if err != nil && !errors.Is(err, context.Canceled) {
		s.logger.Error(op, err)
		shutdownErrors = append(shutdownErrors, err)
	}

	// Проверка на ошибки из notify (если был не interrupt)
	if err, ok := <-s.notify; ok {
		if err != nil && !errors.Is(err, context.Canceled) {
			shutdownErrors = append(shutdownErrors, err)
		}
	}

	return errors.Join(shutdownErrors...)
}