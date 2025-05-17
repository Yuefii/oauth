package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/yuefii/oauth/pkg/helper"
)

var oauthState = "randomstate"

func GithubLoginHandler(w http.ResponseWriter, r *http.Request) {
	conf := GithubOAuthConf()
	url := conf.AuthCodeURL(oauthState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func GithubCallbackHandler(w http.ResponseWriter, r *http.Request) {
	conf := GithubOAuthConf()

	if r.FormValue("state") != oauthState {
		http.Error(w, "invalid state", http.StatusBadRequest)
		return
	}

	token, err := conf.Exchange(context.Background(), r.FormValue("code"))

	if err != nil {
		http.Error(w, "token exchange failed", http.StatusInternalServerError)
		return
	}

	client := conf.Client(context.Background(), token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		http.Error(w, "failed to get user", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var user map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&user)

	username, _ := user["login"].(string)
	fullName, _ := user["name"].(string)
	avatarURL, _ := user["avatar_url"].(string)

	jwtToken, err := helper.GenerateJWT(username, fullName, avatarURL)
	if err != nil {
		http.Error(w, "failed to generate jwt", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"username":  username,
		"full_name": fullName,
		"avatar":    avatarURL,
		"token":     jwtToken,
	})
}
