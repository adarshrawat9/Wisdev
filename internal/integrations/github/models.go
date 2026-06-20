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
	UpdatedAt   time.Time   `json:"updated_at"`
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

type Contributions struct {
	TotalCommits         int       `json:"total_commits"`
	TotalIssues          int       `json:"total_issues"`
	TotalPullRequests    int       `json:"total_pull_requests"`
	TotalReviews         int       `json:"total_reviews"`
	TotalContributions   int       `json:"total_contributions"`
}   

type GrowthAnalytics struct {
	TotalContributions       int     `json:"total_contributions"`
	ActiveDays               int     `json:"active_days"`
	AverageContributionsDay  float64 `json:"average_contributions_per_day"`
	BestDayContributions     int     `json:"best_day_contributions"`
}