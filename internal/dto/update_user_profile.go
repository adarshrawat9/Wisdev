package dto

type UpdateUserProfile struct{
	Bio              *string `json:"bio"`
    GithubUsername   *string `json:"github_username"`
    PortfolioWebsite *string `json:"portfolio_website"`
    AvatarURL        *string `json:"avatar_url"`
}