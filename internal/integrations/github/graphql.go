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

const getUserRepositoriesQuery = `
query($login: String!) {
  user(login: $login) {
    repositories(
      first: 100,
      ownerAffiliations: OWNER,
      orderBy: {
        field: STARGAZERS,
        direction: DESC
      }
    ) {
      nodes {
        name
        description
        stargazerCount
        forkCount

        primaryLanguage {
          name
        }

        url
        createdAt
        updatedAt
      }
    }
  }
}
`