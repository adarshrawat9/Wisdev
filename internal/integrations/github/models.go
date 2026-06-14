package github

type User struct {
	Login       string `json:"login"`
	Name        string `json:"name"`
	Bio         string `json:"bio"`
	AvatarURL   string `json:"avatar_url"`
	Followers   int    `json:"followers"`
	Following   int    `json:"following"`
	PublicRepos int    `json:"public_repos"`
}