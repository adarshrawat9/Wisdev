package github

import "time"

type User struct {
	Name        string
	Bio         string
	AvatarURL   string
	Followers   int
	Following   int
	PublicRepos int
}

type Repository struct {
	Name        string
	Description string
	Stars       int
	Forks       int
	Language    string
	URL         string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}