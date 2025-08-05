package server

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/sgaunet/httpfileserver/internal/config"
	"github.com/sirupsen/logrus"
)

func TestNewServer(t *testing.T) {
	cfg := &config.Config{
		Directory: "/tmp",
		Port:      8080,
		User:      "user",
		Password:  "pass",
	}
	
	logger := logrus.New()
	server := NewServer(cfg, logger)
	
	if server.config != cfg {
		t.Errorf("NewServer() config = %v, want %v", server.config, cfg)
	}
	
	if server.server.Addr != ":8080" {
		t.Errorf("NewServer() server.Addr = %v, want %v", server.server.Addr, ":8080")
	}
	
	if server.server.ReadTimeout != serverReadTimeout {
		t.Errorf("NewServer() server.ReadTimeout = %v, want %v", server.server.ReadTimeout, serverReadTimeout)
	}
	
	if server.server.WriteTimeout != serverWriteTimeout {
		t.Errorf("NewServer() server.WriteTimeout = %v, want %v", server.server.WriteTimeout, serverWriteTimeout)
	}
	
	if server.server.IdleTimeout != serverIdleTimeout {
		t.Errorf("NewServer() server.IdleTimeout = %v, want %v", server.server.IdleTimeout, serverIdleTimeout)
	}
}

func TestServer_setupHandler(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir := t.TempDir()
	
	// Create a test file
	testFile := filepath.Join(tmpDir, "test.txt")
	testContent := "test content"
	err := os.WriteFile(testFile, []byte(testContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	
	tests := []struct {
		name      string
		config    *config.Config
		path      string
		wantAuth  bool
		wantStatus int
	}{
		{
			name: "without auth",
			config: &config.Config{
				Directory: tmpDir,
				Port:      8080,
			},
			path:       "/test.txt",
			wantAuth:   false,
			wantStatus: http.StatusOK,
		},
		{
			name: "with auth - no credentials",
			config: &config.Config{
				Directory: tmpDir,
				Port:      8080,
				User:      "user",
				Password:  "pass",
			},
			path:       "/test.txt",
			wantAuth:   true,
			wantStatus: http.StatusUnauthorized,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := logrus.New()
			server := NewServer(tt.config, logger)
			server.setupHandler()
			
			req := httptest.NewRequest("GET", tt.path, nil)
			rr := httptest.NewRecorder()
			
			server.server.Handler.ServeHTTP(rr, req)
			
			if status := rr.Code; status != tt.wantStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.wantStatus)
			}
			
			if tt.wantAuth && tt.wantStatus == http.StatusUnauthorized {
				authHeader := rr.Header().Get("WWW-Authenticate")
				if authHeader == "" {
					t.Error("expected WWW-Authenticate header for auth-protected endpoint")
				}
			}
		})
	}
}

func TestServer_ListenAndServe(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir := t.TempDir()
	
	// Use a random port to avoid conflicts
	cfg := &config.Config{
		Directory: tmpDir,
		Port:      0, // Let the system assign a port
	}
	
	logger := logrus.New()
	server := NewServer(cfg, logger)
	
	// Start the server in a goroutine
	go func() {
		server.ListenAndServe()
	}()
	
	// Give the server time to start
	time.Sleep(100 * time.Millisecond)
	
	// Since we can't easily test the actual server start without binding issues,
	// we'll just verify that the handler is set up correctly
	if server.server.Handler == nil {
		t.Error("server handler was not set up")
	}
}