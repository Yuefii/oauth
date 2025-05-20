package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"
	"github.com/yuefii/oauth/internal/users"
	"github.com/yuefii/oauth/pkg/helper"
)

var store = sessions.NewCookieStore([]byte("super-secret-key"))

func GithubLoginHandler(w http.ResponseWriter, r *http.Request) {
	state, err := helper.GenerateRandomString(16)
	if err != nil {
		http.Error(w, "failed to generate state", http.StatusInternalServerError)
		return
	}

	session, _ := store.Get(r, "auth-session")
	session.Values["oauth_state"] = state
	session.Save(r, w)

	conf := GithubOAuthConf()
	url := conf.AuthCodeURL(state)
	log.Println("redirecting to GitHub with state:", state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func GithubCallbackHandler(w http.ResponseWriter, r *http.Request) {
	conf := GithubOAuthConf()
	session, _ := store.Get(r, "auth-session")

	expectedState, ok := session.Values["oauth_state"].(string)
	if !ok || r.FormValue("state") != expectedState {
		log.Printf("invalid state: expected %s, got %s\n", expectedState, r.FormValue("state"))
		http.Error(w, "invalid state", http.StatusBadRequest)
		return
	}

	delete(session.Values, "oauth_state")
	session.Save(r, w)

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
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		http.Error(w, "failed to decode GitHub response", http.StatusInternalServerError)
		return
	}

	idFloat, ok := user["id"].(float64)
	if !ok {
		http.Error(w, "invalid GitHub ID format", http.StatusInternalServerError)
		return
	}
	githubID := strconv.FormatInt(int64(idFloat), 10)
	username, _ := user["login"].(string)
	fullName, _ := user["name"].(string)
	avatarURL, _ := user["avatar_url"].(string)

	savedUser, err := users.GetOrCreateUser(githubID, username, fullName, avatarURL)
	if err != nil {
		fmt.Printf("ERROR saving user: %v\n", err)
		http.Error(w, "failed to save user", http.StatusInternalServerError)
		return
	}

	jwtToken, err := helper.GenerateJWT(savedUser.Username, savedUser.FullName, savedUser.AvatarURL)
	if err != nil {
		http.Error(w, "failed to generate jwt", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"username":  savedUser.Username,
		"full_name": savedUser.FullName,
		"avatar":    savedUser.AvatarURL,
		"token":     jwtToken,
	})
}
