// Package logger provides logging functionality for the HTTP file server.
package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// New creates a new logger instance with the specified debug level.
func New(debugLevel string) *logrus.Logger {
	log := logrus.New()
	log.SetOutput(os.Stdout)

	switch debugLevel {
	case "info":
		log.SetLevel(logrus.InfoLevel)
	case "warn":
		log.SetLevel(logrus.WarnLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	default:
		log.SetLevel(logrus.DebugLevel)
	}
	
	return log
}