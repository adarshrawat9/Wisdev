package dto


type UserResponse struct{
	ID               string `json:"id"`
    Username         string `json:"username"`
    Email            string `json:"email"`
    Bio              *string `json:"bio"`
    GithubUsername   *string `json:"github_username"`
    PortfolioWebsite *string `json:"portfolio_website"`
    AvatarURL        *string `json:"avatar_url"`
}