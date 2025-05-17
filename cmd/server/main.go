package main

import (
	"log"
	"net/http"

	"github.com/yuefii/oauth/config"
	"github.com/yuefii/oauth/internal/auth"
)

func main() {
	config.LoadDotEnv()

	http.HandleFunc("/auth/github/login", auth.GithubLoginHandler)
	http.HandleFunc("/auth/github/callback", auth.GithubCallbackHandler)

	log.Println("listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
