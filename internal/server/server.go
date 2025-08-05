// Package server provides HTTP server functionality for the file server.
package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sgaunet/httpfileserver/internal/auth"
	"github.com/sgaunet/httpfileserver/internal/config"
	"github.com/sirupsen/logrus"
)

const (
	serverReadTimeout  = 30 * time.Second
	serverWriteTimeout = 30 * time.Second
	serverIdleTimeout  = 120 * time.Second
)

// Server represents the HTTP file server.
type Server struct {
	config *config.Config
	server *http.Server
	logger *logrus.Logger
}

// NewServer creates a new HTTP server with the given configuration and logger.
func NewServer(cfg *config.Config, logger *logrus.Logger) *Server {
	return &Server{
		config: cfg,
		logger: logger,
		server: &http.Server{
			Addr:         cfg.Address(),
			ReadTimeout:  serverReadTimeout,
			WriteTimeout: serverWriteTimeout,
			IdleTimeout:  serverIdleTimeout,
		},
	}
}

// ListenAndServe starts the HTTP server and serves requests.
func (s *Server) ListenAndServe() error {
	s.setupHandler()
	if err := s.server.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	return nil
}

func (s *Server) setupHandler() {
	fileServer := http.FileServer(http.Dir(s.config.Directory))
	
	if s.config.HasAuth() {
		authenticator := auth.NewAuthenticator(s.config.User, s.config.Password)
		s.server.Handler = authenticator.BasicAuth(fileServer)
		s.logger.Infoln("Launch webserver with basic auth")
	} else {
		s.server.Handler = fileServer
		s.logger.Infoln("Launch webserver without auth")
	}
}