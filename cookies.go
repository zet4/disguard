package disguard // import "go.zeta.pm/disguard"

import (
	"net/http"

	"time"
)

// GetSession ...
func (o *Session) getSession(r *http.Request) (*User, error) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return nil, err
	}

	var user User

	if err = o.cookieHandler.Decode("session", cookie.Value, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// SetSession ...
func (o *Session) setSession(w http.ResponseWriter, user User) error {
	encoded, err := o.cookieHandler.Encode("session", user)
	if err != nil {
		return err
	}

	cookie := &http.Cookie{
		Name:     "session",
		Value:    encoded,
		MaxAge:   60 * 60 * 24 * 7,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)

	return nil
}

// ClearSession ...
func (o *Session) clearSession(w http.ResponseWriter) error {
	cookie := &http.Cookie{
		Name:     "session",
		Value:    "",
		MaxAge:   -1,
		Path:     "/",
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)

	return nil
}
