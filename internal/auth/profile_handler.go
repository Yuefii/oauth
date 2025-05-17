package auth

import (
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(jwt.MapClaims)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"username":  user["username"].(string),
		"full_name": user["fullName"].(string),
		"avatar":    user["avatar"].(string),
	})
}
