package main

import (
	"log"
	"net/http"

	"github.com/yuefii/oauth/config"
	"github.com/yuefii/oauth/internal/auth"
	"github.com/yuefii/oauth/internal/users"
	"github.com/yuefii/oauth/middleware"
)

func main() {
	config.LoadDotEnv()
	config.ConnDB()

	http.HandleFunc("/auth/github/login", auth.GithubLoginHandler)
	http.HandleFunc("/auth/github/callback", auth.GithubCallbackHandler)
	http.Handle("/api/user/profile", middleware.AuthMiddleware(http.HandlerFunc(users.UserProfileHandler)))

	log.Println("listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
