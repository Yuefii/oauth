package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yuefii/oauth/config"
)

func GenerateJWT(username, fullName, avatar string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"fullName": fullName,
		"avatar":   avatar,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(config.JWT_SECRET))
	if err != nil {
		return "", err
	}
	// Debug
	//	log.Println("Generated JWT Token:", signedToken)
	return signedToken, nil
}
