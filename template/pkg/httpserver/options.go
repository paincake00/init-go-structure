package httpserver

import (
	"time"
)

type Option func(s *Server)

func Addr(address string) Option {
	return func(s *Server) {
		s.address = address
	}
}

func Timeouts(readTimeout, writeTimeout, shutdownTimeout time.Duration) Option {
	return func(s *Server) {
		s.readTimeout = readTimeout
		s.writeTimeout = writeTimeout
		s.shutdownTimeout = shutdownTimeout
	}
}
