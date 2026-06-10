package model

import "time"

type User struct {
	ID               string
	Username         string
	Email            string
	PasswordHash     string
	GithubUsername   *string
	PortfolioWebsite *string
	AvatarURL        *string
	Bio              *string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}