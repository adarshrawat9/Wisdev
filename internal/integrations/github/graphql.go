package github


type graphQLRequest struct {
	Query     string         `json:"query"`
	Variables map[string]any `json:"variables"`
}

type graphQLResponse struct {
	Data struct {
		User struct {
			Name      string `json:"name"`
			Bio       string `json:"bio"`
			AvatarURL string `json:"avatarUrl"`

			Followers struct {
				TotalCount int `json:"totalCount"`
			} `json:"followers"`

			Following struct {
				TotalCount int `json:"totalCount"`
			} `json:"following"`

			Repositories struct {
				TotalCount int `json:"totalCount"`
			} `json:"repositories"`
		} `json:"user"`
	} `json:"data"`
}


const getUserQuery = `
query($login: String!) {
  user(login: $login) {
    name
    bio
    avatarUrl

    followers {
      totalCount
    }

    following {
      totalCount
    }

    repositories {
      totalCount
    }
  }
}
`