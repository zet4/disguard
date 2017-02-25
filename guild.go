package disguard // import "go.zeta.pm/disguard"

// Guild mini-container, we only care about ID.
type Guild struct {
	ID string `json:"id"`
	// Name                        string            `json:"name"`
	// Icon                        string            `json:"icon"`
	// Region                      string            `json:"region"`
	// AfkChannelID                string            `json:"afk_channel_id"`
	// EmbedChannelID              string            `json:"embed_channel_id"`
	// OwnerID                     string            `json:"owner_id"`
	// JoinedAt                    Timestamp         `json:"joined_at"`
	// Splash                      string            `json:"splash"`
	// AfkTimeout                  int               `json:"afk_timeout"`
	// MemberCount                 int               `json:"member_count"`
	// VerificationLevel           VerificationLevel `json:"verification_level"`
	// EmbedEnabled                bool              `json:"embed_enabled"`
	// Large                       bool              `json:"large"` // ??
	// DefaultMessageNotifications int               `json:"default_message_notifications"`
	// Roles                       []*Role           `json:"roles"`
	// Emojis                      []*Emoji          `json:"emojis"`
	// Members                     []*Member         `json:"members"`
	// Presences                   []*Presence       `json:"presences"`
	// Channels                    []*Channel        `json:"channels"`
	// VoiceStates                 []*VoiceState     `json:"voice_states"`
	// Unavailable                 bool              `json:"unavailable"`
}
