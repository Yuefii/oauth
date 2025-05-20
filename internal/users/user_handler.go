package users

import (
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func UserProfileHandler(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("user").(jwt.MapClaims)
	if !ok {
		http.Error(w, "unautorized", http.StatusUnauthorized)
		return
	}

	username, _ := claims["username"].(string)
	fullName, _ := claims["fullName"].(string)
	avatar, _ := claims["avatar"].(string)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"username":  username,
		"full_name": fullName,
		"avatar":    avatar,
	})
}
