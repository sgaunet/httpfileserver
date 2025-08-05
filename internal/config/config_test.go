package config

import (
	"os"
	"testing"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name      string
		dir       string
		port      int
		wantErr   bool
		errMsg    string
		setupEnv  map[string]string
		cleanupEnv []string
	}{
		{
			name:    "valid config without auth",
			dir:     "/tmp",
			port:    8080,
			wantErr: false,
		},
		{
			name:    "empty directory",
			dir:     "",
			port:    8080,
			wantErr: true,
			errMsg:  "directory cannot be empty",
		},
		{
			name:    "port below minimum",
			dir:     "/tmp",
			port:    1023,
			wantErr: true,
			errMsg:  "port cannot be under minimum: 1024",
		},
		{
			name:    "port at minimum",
			dir:     "/tmp",
			port:    1024,
			wantErr: false,
		},
		{
			name:    "valid config with auth",
			dir:     "/tmp",
			port:    8080,
			wantErr: false,
			setupEnv: map[string]string{
				"HTTP_USER":     "testuser",
				"HTTP_PASSWORD": "testpass",
			},
			cleanupEnv: []string{"HTTP_USER", "HTTP_PASSWORD"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup environment
			for k, v := range tt.setupEnv {
				os.Setenv(k, v)
			}
			defer func() {
				for _, k := range tt.cleanupEnv {
					os.Unsetenv(k)
				}
			}()

			config, err := NewConfig(tt.dir, tt.port)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Error() != tt.errMsg {
				t.Errorf("NewConfig() error = %v, want %v", err, tt.errMsg)
				return
			}
			if err == nil {
				if config.Directory != tt.dir {
					t.Errorf("Config.Directory = %v, want %v", config.Directory, tt.dir)
				}
				if config.Port != tt.port {
					t.Errorf("Config.Port = %v, want %v", config.Port, tt.port)
				}
			}
		})
	}
}

func TestConfig_HasAuth(t *testing.T) {
	tests := []struct {
		name     string
		user     string
		password string
		want     bool
	}{
		{
			name:     "both user and password",
			user:     "user",
			password: "pass",
			want:     true,
		},
		{
			name:     "only user",
			user:     "user",
			password: "",
			want:     false,
		},
		{
			name:     "only password",
			user:     "",
			password: "pass",
			want:     false,
		},
		{
			name:     "neither user nor password",
			user:     "",
			password: "",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				User:     tt.user,
				Password: tt.password,
			}
			if got := c.HasAuth(); got != tt.want {
				t.Errorf("Config.HasAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_Address(t *testing.T) {
	tests := []struct {
		name string
		port int
		want string
	}{
		{
			name: "default port",
			port: 8081,
			want: ":8081",
		},
		{
			name: "custom port",
			port: 3000,
			want: ":3000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Port: tt.port,
			}
			if got := c.Address(); got != tt.want {
				t.Errorf("Config.Address() = %v, want %v", got, tt.want)
			}
		})
	}
}