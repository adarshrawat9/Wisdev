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
	Name             string         `json:"name"`
	Description      string         `json:"description"`

	Language         string         `json:"language"`

	Stars            int 			`json:"stargazers_count"`
	Forks            int 			`json:"forks_count"`

	HTMLURL          string 		`json:"html_url"`

	Topics           []string 		`json:"topics"`

	OpenIssues       int 			`json:"open_issues_count"`

	Fork             bool 			`json:"fork"`

	CreatedAt        time.Time 		`json:"created_at"`
	UpdatedAt        time.Time 		`json:"updated_at"`
}        