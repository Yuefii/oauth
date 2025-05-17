package auth

import (
	"github.com/yuefii/oauth/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func GithubOAuthConf() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     config.GetDotEnv("GITHUB_CLIENT_ID", ""),
		ClientSecret: config.GetDotEnv("GITHUB_CLIENT_SECRET", ""),
		RedirectURL:  config.GetDotEnv("GITHUB_REDIRECT_URL", ""),

		Scopes:   []string{"user:email"},
		Endpoint: github.Endpoint,
	}
}
