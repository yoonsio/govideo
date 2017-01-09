package govideo

import "net/http"

// AuthClient provides authentication
type AuthClient struct {
}

// User object that persists into database
type User struct {
}

// NewAuthClient creates new AuthClient with random key
func NewAuthClient() *AuthClient {
	return &AuthClient{}
}

// GetUser returns User object from session
func (ac AuthClient) GetUser() *User {
	return nil
}

// SetUser queries database for existing user and stores in session
func (ac AuthClient) SetUser() error {
	// Store user into session (securecookie) instead of context
	return nil
}

// ClearUser removes authentication session from cookie
func (ac AuthClient) ClearUser() {
}

// AuthMiddleware for authentication
func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: check session(securecookie)
		// TODO: verify authentication
		h.ServeHTTP(w, r)
	})
}
