package govideo

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/sickyoon/govideo/govideo/models"
)

// AuthClient provides authentication
type AuthClient struct {
	store *sessions.CookieStore // cookie session to store user info
	db    *MongoClient          // db client to store user info
}

// NewAuthClient creates new AuthClient with random key
func NewAuthClient(store *sessions.CookieStore, db *MongoClient) *AuthClient {
	return &AuthClient{store, db}
}

// SetUser queries database for existing user and stores in session
func (ac AuthClient) SetUser() error {
	// Store user into session (securecookie) instead of context
	return nil
}

// GetUser returns User object from session
func (ac AuthClient) GetUser(username, pass string) (*models.User, error) {
	// TODO: hash password
	// TODO: get user info from DB
	user, err := ac.db.GetUser("", "")
	if err != nil {
		return nil, err
	}
	// TODO: generate unique cache key associated with user
	// TODO: set cache using freecache
	// TODO: store unique cache key into session cookie
	return user, nil
}

// ClearUser removes authentication session from cookie
func (ac AuthClient) ClearUser() {
}

// AuthMiddleware for authentication
func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: get cache_key from session
		// TODO: validate user
		// TODO: if not, redirect to login
		if true {
			http.Redirect(w, r, "/login", http.StatusFound)
		}

		h.ServeHTTP(w, r)
	})
}
