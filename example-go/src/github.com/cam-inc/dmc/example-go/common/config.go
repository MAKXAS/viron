package common

import (
	"fmt"
	"os"
)

// Config of
type Config struct {
	Scheme string
	Host   string
	Port   int16
	GoogleOAuth
	DefaultRole string
}

// GoogleOAuth of
type GoogleOAuth struct {
	ClientID          string
	ClientSecret      string
	RedirectURL       string
	Scopes            []string
	AllowEmailDomains []string
}

var config *Config

func init() {
	config = &Config{
		Scheme:      "http",
		Host:        "localhost",
		Port:        3000,
		DefaultRole: "viewer", // AdminUser作成時の初期ロール

		GoogleOAuth: GoogleOAuth{
			ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
			RedirectURL:  "http://localhost:3000/googleoauth2callback",
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
			},
			AllowEmailDomains: []string{
				// ここに書いたドメインがGoogle認証で利用可能
				"camobile.com",
			},
		},
	}
}

// GetHostName of
func GetHostName() string {
	return fmt.Sprintf("%s://%s:%d", config.Scheme, config.Host, config.Port)
}

// GetGoogleOAuth of
func GetGoogleOAuth() GoogleOAuth {
	return config.GoogleOAuth
}

// GetDefaultRole of
func GetDefaultRole() string {
	return config.DefaultRole
}
