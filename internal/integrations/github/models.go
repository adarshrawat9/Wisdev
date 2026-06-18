package github

import "time"

type User struct {
	Name        string `json:"name"`
	Bio         string `json:"bio"`
	AvatarURL   string `json:"avatar_url"`
	Followers   int    `json:"followers"`
	Following   int    `json:"following"`
	PublicRepos int    `json:"public_repos"`
}

type Repository struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Stars       int         `json:"stars"`
	Forks       int         `json:"forks"`
	Language    string      `json:"language"`
	URL         string      `json:"url"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at`
}

type TopRepository struct {
	Name  string       `json:"name"`
	Stars int          `json:"stars"`
	URL   string       `json:"url"`
}

type LanguageStats struct {
	Name         string `json:"name"`
	Repositories int    `json:"repositories"`
	Stars        int    `json:"stars"`
	Forks        int    `json:"forks"`
}

type Analytics struct {
	TotalRepositories int                `json:"total_repositories"`
	TotalStars        int                `json:"total_stars"`
	TotalForks        int                `json:"total_forks"`
	MostUsedLanguage  string             `json:"most_used_language"`
	Languages         []LanguageStats    `json:"top_language"`
	TopRepositories   []TopRepository    `json:"top_repositories"`
}      