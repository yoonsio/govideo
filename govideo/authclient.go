package govideo

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/sickyoon/govideo/govideo/models"
)

// ContextKey -
type ContextKey string

// AuthClient provides authentication
type AuthClient struct {
	*sessions.CookieStore // cookie session to store user info
	*MongoClient          // db client to store user info
	*RedisClient
	redirectURI string
	cookieKey   string
	sessionKey  string
	contextKey  ContextKey
}

// NewAuthClient creates new AuthClient with random key
func NewAuthClient(store *sessions.CookieStore, db *MongoClient, cache *RedisClient) *AuthClient {
	return &AuthClient{
		CookieStore: store,
		MongoClient: db,
		RedisClient: cache,
		redirectURI: "/login",
		cookieKey:   "auth", // TODO: random-generated string key
		sessionKey:  "user",
		contextKey:  "auth",
	}
}

// AuthMiddleware for authentication
func (ac AuthClient) AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := ac.CurUser(r)
		if err != nil {
			ac.Redirect(w, r)
			return
		}
		ctx := context.WithValue(r.Context(), ac.contextKey, user)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Redirect redirects to authenticate uri
func (ac AuthClient) Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, ac.redirectURI, http.StatusFound)
}

// Authenticate authenticates user with provided credentials
// To be called by login handler
func (ac *AuthClient) Authenticate(w http.ResponseWriter, r *http.Request, email, password string) (*models.User, error) {
	dbUser, err := ac.validate(email, password)
	if err != nil {
		return nil, err
	}
	cacheUser, err := ac.setCache(dbUser)
	if err != nil {
		return nil, err
	}
	err = ac.setSession(w, r, cacheUser)
	return dbUser, err
}

// CurUser gets currently logged-in user
func (ac *AuthClient) CurUser(r *http.Request) (*models.User, error) {
	key, err := ac.getSession(r)
	if err != nil {
		return nil, err
	}
	user, err := ac.getCache(key)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// ClearUser removes current user session
func (ac *AuthClient) ClearUser(w http.ResponseWriter, r *http.Request) error {
	key, err := ac.getSession(r)
	if err != nil {
		return err
	}
	err = ac.clearSession(w, r)
	if err != nil {
		return err
	}
	return ac.clearCache(key)
}

// Validate authenticates user against database
func (ac *AuthClient) validate(email, password string) (*models.User, error) {
	// TODO: validate email/password
	return ac.GetUserFromDB(email, password)
}

func (ac *AuthClient) setCache(user *models.User) ([]byte, error) {

	// serialize user
	userBytes, err := user.Marshal()
	if err != nil {
		return nil, err
	}

	// set user to redis cache
	return ac.SetAuthCache(user.Email, userBytes)
}

func (ac *AuthClient) getCache(key []byte) (*models.User, error) {

	// get user from redis cache
	userBytes, err := ac.GetAuthCache(key)
	if err != nil {
		return nil, err
	}

	// deserialize user
	user := models.User{}
	err = user.Unmarshal(userBytes)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ac *AuthClient) clearCache(key []byte) error {
	return ac.ClearAuthCache(key)
}

func (ac *AuthClient) setSession(w http.ResponseWriter, r *http.Request, user []byte) error {
	session, err := ac.Get(r, ac.cookieKey)
	if err != nil {
		return err
	}
	session.Values[ac.sessionKey] = user
	return session.Save(r, w)
}

func (ac *AuthClient) clearSession(w http.ResponseWriter, r *http.Request) error {
	session, err := ac.Get(r, ac.cookieKey)
	if err != nil {
		return err
	}
	if _, ok := session.Values[ac.sessionKey]; ok {
		delete(session.Values, ac.sessionKey)
	}
	return session.Save(r, w)
}

func (ac *AuthClient) getSession(r *http.Request) ([]byte, error) {
	session, err := ac.Get(r, ac.cookieKey)
	if err != nil {
		return nil, nil
	}
	userData, ok := session.Values[ac.sessionKey]
	if !ok {
		return nil, fmt.Errorf("failed to get user session")
	}
	return userData.([]byte), nil
}
