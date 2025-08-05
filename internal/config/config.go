// Package config provides configuration management for the HTTP file server.
package config

import (
	"errors"
	"fmt"
	"os"
)

var (
	// ErrEmptyDirectory is returned when no directory is specified.
	ErrEmptyDirectory = errors.New("directory cannot be empty")
	// ErrPortTooLow is returned when the port is below the minimum allowed value.
	ErrPortTooLow     = errors.New("port cannot be under minimum")
)

const (
	// DefaultPort is the default port used by the server.
	DefaultPort = 8081
	minPort     = 1024
)

// Config holds the server configuration.
type Config struct {
	Directory string
	Port      int
	User      string
	Password  string
}

// NewConfig creates a new configuration with validation.
func NewConfig(dir string, port int) (*Config, error) {
	if port == 0 {
		port = DefaultPort
	}
	
	if dir == "" {
		return nil, ErrEmptyDirectory
	}

	if port < minPort {
		return nil, fmt.Errorf("%w: %d", ErrPortTooLow, minPort)
	}

	return &Config{
		Directory: dir,
		Port:      port,
		User:      os.Getenv("HTTP_USER"),
		Password:  os.Getenv("HTTP_PASSWORD"),
	}, nil
}

// HasAuth returns true if both username and password are configured.
func (c *Config) HasAuth() bool {
	return c.User != "" && c.Password != ""
}

// Address returns the server address in the format ":port".
func (c *Config) Address() string {
	return fmt.Sprintf(":%d", c.Port)
}