package disguard // import "go.zeta.pm/disguard"

import (
	"fmt"
	"net/url"
)

var (
	oauthScope       = url.QueryEscape("identify guilds")
	authorizationURL = "https://discordapp.com/oauth2/authorize?response_type=code&client_id=%s&scope=%s&redirect_uri=%s"
	tokenURL         = "https://discordapp.com/api/oauth2/token?client_id=%s&client_secret=%s&grant_type=authorization_code&scope=%s&code=%s&redirect_uri=%s"
)

// TokenResponse contains response from OAuth
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

func (conf *OAuthSection) getRedirectURL() string {
	if conf.redirectURL != nil {
		return *conf.redirectURL
	}
	temp := url.QueryEscape(conf.RedirectURL)
	conf.redirectURL = &temp
	return temp
}

// GetAuthorizationURL returns blah
func (conf *OAuthSection) GetAuthorizationURL() string {
	if conf.authorizationURL != nil {
		return *conf.authorizationURL
	}

	temp := fmt.Sprintf(authorizationURL, conf.ClientID, oauthScope, conf.getRedirectURL())
	conf.authorizationURL = &temp
	return temp
}

// GetTokenURL returns blah
func (conf *OAuthSection) GetTokenURL(code string) string {
	return fmt.Sprintf(tokenURL, conf.ClientID, conf.ClientSecret, oauthScope, code, conf.getRedirectURL())
}
