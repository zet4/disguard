package disguard // import "go.zeta.pm/disguard"

import (
	"fmt"
	"net/url"
)

var oauthScope = url.QueryEscape("identify guilds")

// OAuthConfig contains info for OAuth
type OAuthConfig struct {
	ClientID         string
	ClientSecret     string
	RedirectURL      string
	AuthorizationURL string
	TokenURL         string

	authorizationURL, redirectURL *string
}

// TokenResponse contains response from OAuth
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

func (conf *OAuthConfig) getRedirectURL() string {
	if conf.redirectURL != nil {
		return *conf.redirectURL
	}
	temp := url.QueryEscape(conf.RedirectURL)
	conf.redirectURL = &temp
	return temp
}

// GetAuthorizationURL returns blah
func (conf *OAuthConfig) GetAuthorizationURL() string {
	if conf.authorizationURL != nil {
		return *conf.authorizationURL
	}

	temp := fmt.Sprintf(conf.AuthorizationURL, conf.ClientID, oauthScope, conf.getRedirectURL())
	conf.authorizationURL = &temp
	return temp
}

// GetTokenURL returns blah
func (conf *OAuthConfig) GetTokenURL(code string) string {
	return fmt.Sprintf(conf.TokenURL, conf.ClientID, conf.ClientSecret, oauthScope, code, conf.getRedirectURL())
}

// NewOAuth returns an OAuthConfig using viper.
func NewOAuth(o OAuthSection) *OAuthConfig {
	conf := &OAuthConfig{
		ClientID:         o.ClientID,
		ClientSecret:     o.ClientSecret,
		RedirectURL:      o.RedirectURL,
		AuthorizationURL: "https://discordapp.com/oauth2/authorize?response_type=code&client_id=%s&scope=%s&redirect_uri=%s",
		TokenURL:         "https://discordapp.com/api/oauth2/token?client_id=%s&client_secret=%s&grant_type=authorization_code&scope=%s&code=%s&redirect_uri=%s",
	}
	temp1 := url.QueryEscape(conf.RedirectURL)
	conf.redirectURL = &temp1

	temp2 := fmt.Sprintf(conf.AuthorizationURL, conf.ClientID, oauthScope, conf.getRedirectURL())
	conf.authorizationURL = &temp2
	return conf
}
