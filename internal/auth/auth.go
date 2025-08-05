// Package auth provides HTTP Basic Authentication middleware.
package auth

import (
	"crypto/sha256"
	"crypto/subtle"
	"net/http"
)

// Authenticator handles HTTP Basic Authentication using SHA256 hashing.
type Authenticator struct {
	User     string
	Password string
}

// NewAuthenticator creates a new authenticator with the given credentials.
func NewAuthenticator(user, password string) *Authenticator {
	return &Authenticator{
		User:     user,
		Password: password,
	}
}

// BasicAuth provides HTTP Basic Authentication middleware with constant-time comparison.
func (a *Authenticator) BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok {
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))
			expectedUsernameHash := sha256.Sum256([]byte(a.User))
			expectedPasswordHash := sha256.Sum256([]byte(a.Password))

			usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
			passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

			if usernameMatch && passwordMatch {
				next.ServeHTTP(w, r)
				return
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}