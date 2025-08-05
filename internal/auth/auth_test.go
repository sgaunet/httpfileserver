package auth

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthenticator_BasicAuth(t *testing.T) {
	auth := NewAuthenticator("testuser", "testpass")
	
	// Create a simple handler to protect
	protectedHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("protected content"))
	})
	
	// Wrap the handler with basic auth
	handler := auth.BasicAuth(protectedHandler)
	
	tests := []struct {
		name       string
		authHeader string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "no auth header",
			authHeader: "",
			wantStatus: http.StatusUnauthorized,
			wantBody:   "Unauthorized\n",
		},
		{
			name:       "correct credentials",
			authHeader: "Basic " + base64.StdEncoding.EncodeToString([]byte("testuser:testpass")),
			wantStatus: http.StatusOK,
			wantBody:   "protected content",
		},
		{
			name:       "wrong username",
			authHeader: "Basic " + base64.StdEncoding.EncodeToString([]byte("wronguser:testpass")),
			wantStatus: http.StatusUnauthorized,
			wantBody:   "Unauthorized\n",
		},
		{
			name:       "wrong password",
			authHeader: "Basic " + base64.StdEncoding.EncodeToString([]byte("testuser:wrongpass")),
			wantStatus: http.StatusUnauthorized,
			wantBody:   "Unauthorized\n",
		},
		{
			name:       "malformed auth header",
			authHeader: "Basic malformed",
			wantStatus: http.StatusUnauthorized,
			wantBody:   "Unauthorized\n",
		},
		{
			name:       "empty credentials",
			authHeader: "Basic " + base64.StdEncoding.EncodeToString([]byte(":")),
			wantStatus: http.StatusUnauthorized,
			wantBody:   "Unauthorized\n",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			if tt.authHeader != "" {
				req.Header.Set("Authorization", tt.authHeader)
			}
			
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)
			
			if status := rr.Code; status != tt.wantStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.wantStatus)
			}
			
			if body := rr.Body.String(); body != tt.wantBody {
				t.Errorf("handler returned unexpected body: got %v want %v",
					body, tt.wantBody)
			}
			
			// Check WWW-Authenticate header for unauthorized requests
			if tt.wantStatus == http.StatusUnauthorized {
				authHeader := rr.Header().Get("WWW-Authenticate")
				expectedHeader := `Basic realm="restricted", charset="UTF-8"`
				if authHeader != expectedHeader {
					t.Errorf("handler returned wrong WWW-Authenticate header: got %v want %v",
						authHeader, expectedHeader)
				}
			}
		})
	}
}

func TestNewAuthenticator(t *testing.T) {
	user := "testuser"
	password := "testpass"
	
	auth := NewAuthenticator(user, password)
	
	if auth.User != user {
		t.Errorf("NewAuthenticator() User = %v, want %v", auth.User, user)
	}
	
	if auth.Password != password {
		t.Errorf("NewAuthenticator() Password = %v, want %v", auth.Password, password)
	}
}