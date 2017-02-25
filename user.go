package main // import "go.zeta.pm/disguard"

type User struct {
	Avatar        string   `json:"avatar"`
	Discriminator string   `json:"discriminator"`
	ID            string   `json:"id"`
	Name          string   `json:"username"`
	Whitelisted   []string `json:"whitelisted"`
}
