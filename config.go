package disguard // import "go.zeta.pm/disguard"

type (
	// Config ...
	Config struct {
		ListenAddress string `yaml:"listen_address"`
		ProxyAddress  string `yaml:"proxy_address"`

		HeaderName        string   `yaml:"header_name"`
		WhitelistedGuilds []string `yaml:"whitelisted_guilds"`
		RequireSession    bool     `yaml:"require_session"`
		IgnoredPaths      []string `yaml:"ignored_paths"`
		AuthRoot          string   `yaml:"auth_root"`

		OAuth   OAuthSection   `yaml:"oauth"`
		Session SessionSection `yaml:"session"`
	}

	// OAuthSection ...
	OAuthSection struct {
		RedirectURL  string `yaml:"redirect_url"`
		ClientID     string `yaml:"client_id"`
		ClientSecret string `yaml:"client_secret"`

		redirectURL, authorizationURL *string
	}

	// SessionSection ...
	SessionSection struct {
		HashKey  string `yaml:"hash_key"`
		BlockKey string `yaml:"block_key"`
	}
)
