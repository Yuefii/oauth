package users

import (
	"database/sql"
	"fmt"

	"github.com/yuefii/oauth/config"
	"github.com/yuefii/oauth/models"
)

func GetOrCreateUser(githubID, username, fullName, avatarURL string) (*models.User, error) {
	if config.DB == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	db := config.DB
	var user models.User

	err := db.QueryRow(`SELECT id, github_id, username, full_name, avatar_url FROM users WHERE github_id = ?`, githubID).Scan(&user.ID, &user.GithubID, &user.Username, &user.FullName, &user.AvatarURL)

	if err == sql.ErrNoRows {
		res, err := db.Exec(`INSERT INTO users (github_id, username, full_name, avatar_url) VALUES (?,?,?,?)`, githubID, username, fullName, avatarURL)
		if err != nil {
			return nil, err
		}

		id, err := res.LastInsertId()
		user = models.User{
			ID:        id,
			GithubID:  githubID,
			Username:  username,
			FullName:  fullName,
			AvatarURL: avatarURL,
		}
		return &user, nil
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}
