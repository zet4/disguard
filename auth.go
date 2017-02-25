package main // import "go.zeta.pm/disguard"

import (
	"encoding/json"
	"mime"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/securecookie"
	"github.com/pressly/chi"
)

// Session contains everything OAuth
type Session struct {
	config        *Config
	oauthConfig   *OAuthConfig
	httpClient    *http.Client
	cookieHandler *securecookie.SecureCookie
	box           *rice.Box
}

// HandleLogin ...
func (o *Session) HandleLogin(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, o.oauthConfig.GetAuthorizationURL(), http.StatusFound)
}

// HandleLogout ...
func (o *Session) HandleLogout(w http.ResponseWriter, r *http.Request) {
	o.ClearSession(w)

	http.Redirect(w, r, "/", http.StatusFound)
}

// ObtainToken ...
func (o *Session) ObtainToken(code string) (*TokenResponse, error) {
	req, err := http.NewRequest("POST", o.oauthConfig.GetTokenURL(code), nil)
	if err != nil {
		return nil, err
	}
	resp, err := o.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, http.ErrAbortHandler
	}
	var tokenResponse TokenResponse
	err = json.NewDecoder(resp.Body).Decode(&tokenResponse)
	if err != nil {
		return nil, err
	}
	if !strings.Contains(tokenResponse.Scope, "identify") {
		return nil, http.ErrAbortHandler
	}
	if !strings.Contains(tokenResponse.Scope, "guilds") {
		return nil, http.ErrAbortHandler
	}
	if tokenResponse.TokenType != "Bearer" {
		return nil, http.ErrAbortHandler
	}
	return &tokenResponse, nil
}

// GetUser ...
func (o *Session) GetUser(tokenResponse *TokenResponse) (*User, error) {
	req, err := http.NewRequest("GET", "https://discordapp.com/api/users/@me", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+tokenResponse.AccessToken)
	resp, err := o.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, http.ErrAbortHandler
	}
	var userResponse User
	err = json.NewDecoder(resp.Body).Decode(&userResponse)
	if err != nil {
		return nil, err
	}
	return &userResponse, nil
}

// GetGuilds ...
func (o *Session) GetGuilds(tokenResponse *TokenResponse) ([]Guild, error) {
	req, err := http.NewRequest("GET", "https://discordapp.com/api/users/@me/guilds", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+tokenResponse.AccessToken)
	resp, err := o.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var guildsResponse []Guild
	err = json.NewDecoder(resp.Body).Decode(&guildsResponse)
	if err != nil {
		return nil, http.ErrAbortHandler
	}
	return guildsResponse, nil
}

// FilterWhitelisted itereates over all guilds and returns list of matching guilds
func (o *Session) FilterWhitelisted(guilds []Guild) []string {
	results := make([]string, 0)
mainLabel:
	for _, guild := range guilds {
		for _, whitelistedID := range o.config.WhitelistedGuilds {
			if guild.ID == whitelistedID {
				results = append(results, whitelistedID)
			}
			if len(results) >= len(o.config.WhitelistedGuilds) {
				// Found all the whitelisted guilds, no more need, breaking out.
				break mainLabel
			}
		}
	}
	return results
}

// DoCallback ...
func (o *Session) DoCallback(code string) (*User, error) {
	tokenResponse, err := o.ObtainToken(code)
	if err != nil {
		return nil, err
	}
	userResponse, err := o.GetUser(tokenResponse)
	if err != nil {
		return nil, err
	}
	guildsResponse, err := o.GetGuilds(tokenResponse)
	if err != nil {
		return nil, err
	}
	whitelisted := o.FilterWhitelisted(guildsResponse)
	if len(whitelisted) == 0 {
		return nil, http.ErrAbortHandler
	}
	userResponse.Whitelisted = whitelisted

	return userResponse, nil
}

// HandleCallback ...
func (o *Session) HandleCallback(w http.ResponseWriter, r *http.Request) {
	code, ok := r.URL.Query()["code"]
	if !ok {
		w.WriteHeader(403)
		w.Write(o.box.MustBytes("error403.html"))
		return
	}

	userResponse, err := o.DoCallback(code[0])
	if err != nil {
		w.WriteHeader(403)
		w.Write(o.box.MustBytes("error403.html"))
		return
	}

	err = o.SetSession(w, *userResponse)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

// Route ...
func (o *Session) Route(r chi.Router) {
	r.Get("/login", o.HandleLogin)
	r.Get("/logout", o.HandleLogout)
	r.Get("/callback", o.HandleCallback)

	r.Mount("/static/", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		path := strings.SplitN(r.URL.Path, "/static/", 2)[1]
		if b, err := o.box.Bytes(path); err == nil {
			rw.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(path)))
			rw.Write(b)
		} else {
			rw.WriteHeader(404)
		}
	}))
}

// NewSessionRouter returns handler for session stuff
func NewSessionRouter(c *Config) *Session {
	sess := &Session{
		config:        c,
		oauthConfig:   NewOAuth(c.OAuth),
		httpClient:    &http.Client{},
		cookieHandler: securecookie.New([]byte(c.Session.HashKey), []byte(c.Session.BlockKey)),
		box:           rice.MustFindBox("static"),
	}

	return sess
}
