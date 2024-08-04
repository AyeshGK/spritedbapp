package sessions

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

var sessionName = "db-admin-session"

// InitializeSession initializes the session store with options.
func InitializeSession() {
	// Customize session options if needed
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600, // Session expires in 1 hour
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   false, // Set true in production with HTTPS
	}
}

// SetSession sets a session value for a given key.
func SetSession(w http.ResponseWriter, r *http.Request, key string, value interface{}) {
	session, err := store.Get(r, sessionName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values[key] = value
	session.Save(r, w)
}

// GetSession retrieves a session value for a given key.
func GetSession(r *http.Request, key string) interface{} {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return nil
	}
	return session.Values[key]
}

// ClearSession clears the session.
func ClearSession(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, sessionName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Options.MaxAge = -1 // Clear session by setting MaxAge to -1
	session.Save(r, w)
}

// IsAuthenticated checks if the user is authenticated.
func IsAuthenticated(r *http.Request) bool {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return false
	}
	authenticated := session.Values["authenticated"]
	if authenticated != nil && authenticated.(bool) {
		return true
	}
	return false
}
