package logger

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name      string
		debugLevel string
		wantLevel logrus.Level
	}{
		{
			name:      "info level",
			debugLevel: "info",
			wantLevel: logrus.InfoLevel,
		},
		{
			name:      "warn level",
			debugLevel: "warn",
			wantLevel: logrus.WarnLevel,
		},
		{
			name:      "error level",
			debugLevel: "error",
			wantLevel: logrus.ErrorLevel,
		},
		{
			name:      "default to debug",
			debugLevel: "unknown",
			wantLevel: logrus.DebugLevel,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log := New(tt.debugLevel)
			
			if log.GetLevel() != tt.wantLevel {
				t.Errorf("New() level = %v, want %v", log.GetLevel(), tt.wantLevel)
			}
		})
	}
}