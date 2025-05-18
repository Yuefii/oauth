package models

type User struct {
	ID        int64
	GithubID  string
	Username  string
	FullName  string
	AvatarURL string
}
