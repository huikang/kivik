package auth

import (
	"encoding/json"
	"net/http"

	"github.com/flimzy/kivik/authdb"
)

// Handler is an auth handler.
type Handler interface {
	// MethodName identifies the handler.
	MethodName() string
	// Authenticate authenticates the HTTP request against the user store.
	// On success, a user context must be returned. Any error other than
	// Unauthorized will immediately terminate the authentication process,
	// returning an error to the client.
	Authenticate(*http.Request, authdb.UserStore) (*authdb.UserContext, error)
}

// Session represents an authenticated session.
type Session struct {
	AuthMethod string
	Handlers   []string
	User       *authdb.UserContext
}

// MarshalJSON satisfies the json.Marshaler interface.
func (s *Session) MarshalJSON() ([]byte, error) {
	user := s.User
	if user == nil {
		user = &authdb.UserContext{}
	}
	result := map[string]interface{}{
		"info": map[string]interface{}{
			"authenticated":           s.AuthMethod,
			"authentication_db":       user.AuthDatabase,
			"authentication_handlers": s.Handlers,
		},
		"ok":      true,
		"userCtx": user,
	}
	return json.Marshal(result)
}
