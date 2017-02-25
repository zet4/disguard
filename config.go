package disguard // import "go.zeta.pm/disguard"

type (
	// Config ...
	Config struct {
		ListenAddress string `yaml:"listen_address"`
		ProxyAddress  string `yaml:"proxy_address"`

		HeaderName        string   `yaml:"header_name"`
		WhitelistedGuilds []string `yaml:"whitelisted_guilds"`
		RequireSession    bool     `yaml:"require_session"`

		OAuth   OAuthSection   `yaml:"oauth"`
		Session SessionSection `yaml:"session"`
	}

	// OAuthSection ...
	OAuthSection struct {
		RedirectURL  string `yaml:"redirect_url"`
		ClientID     string `yaml:"client_id"`
		ClientSecret string `yaml:"client_secret"`
	}

	// SessionSection ...
	SessionSection struct {
		HashKey  string `yaml:"hash_key"`
		BlockKey string `yaml:"block_key"`
	}
)
