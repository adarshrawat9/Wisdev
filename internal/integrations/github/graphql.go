package github

import "time"

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


type repositoriesResponse struct {
	Data struct {
		User struct {
			Repositories struct {
				Nodes []struct {
					Name           string `json:"name"`
					Description    string `json:"description"`
					StargazerCount int    `json:"stargazerCount"`
					ForkCount      int    `json:"forkCount"`

					PrimaryLanguage *struct {
						Name string `json:"name"`
					} `json:"primaryLanguage"`

					URL       string    `json:"url"`
					CreatedAt time.Time `json:"createdAt"`
					UpdatedAt time.Time `json:"updatedAt"`
				} `json:"nodes"`
			} `json:"repositories"`
		} `json:"user"`
	} `json:"data"`
}

type contributionsResponse struct {
	Data struct {
		User struct {
			ContributionsCollection struct {
				TotalCommitContributions            int `json:"totalCommitContributions"`
				TotalIssueContributions             int `json:"totalIssueContributions"`
				TotalPullRequestContributions       int `json:"totalPullRequestContributions"`
				TotalPullRequestReviewContributions int `json:"totalPullRequestReviewContributions"`
			} `json:"contributionsCollection"`
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

const getUserContributionsQuery = `
query($login: String!) {
  user(login: $login) {

    contributionsCollection {

      totalCommitContributions

      totalIssueContributions

      totalPullRequestContributions

      totalPullRequestReviewContributions
    }
  }
}
`