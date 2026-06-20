package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Client struct{
	httpClient *http.Client
	token      string
}


func NewClient() *Client{
	return &Client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		token: os.Getenv("GITHUB_TOKEN"),
	}
}



func (c *Client) GetUser(username string) (*User, error){

	reqBody := graphQLRequest{
		Query: getUserQuery,
		Variables: map[string]any{
			"login" : username,
		},
	}

	body, err := json.Marshal(reqBody)
	if err != nil{
		return nil, err
	}
	
	req, err := http.NewRequest(
		http.MethodPost,
	    "https://api.github.com/graphql",
	    bytes.NewBuffer(body),
	)
	if err != nil{
		return nil, err
	}

	req.Header.Set(
		"Authorization",
		"Bearer " + c.token,
	)
	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	response, err := c.httpClient.Do(req)
	if err != nil{
		return nil, err
	}
	defer response.Body.Close()

	var gqlResp graphQLResponse

err = json.NewDecoder(response.Body).Decode(&gqlResp)
if err != nil {
	return nil, err
}

user := &User{
	Name: gqlResp.Data.User.Name,
	Bio: gqlResp.Data.User.Bio,
	AvatarURL: gqlResp.Data.User.AvatarURL,
	Followers: gqlResp.Data.User.Followers.TotalCount,
	Following: gqlResp.Data.User.Following.TotalCount,
	PublicRepos: gqlResp.Data.User.Repositories.TotalCount,
}

return user, nil
}


func (c *Client) GetUserRepositories(username string) ([]Repository, error){

	reqBody := graphQLRequest{
		Query: getUserRepositoriesQuery,
		Variables: map[string]any{
			"login" : username,
		},
	}

	body, err := json.Marshal(reqBody)
	if err != nil{
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		"https://api.github.com/graphql",
		bytes.NewBuffer(body),
	)
	if err != nil{
		return nil, err
	}

	req.Header.Set(
		"Authorization",
		"Bearer " + c.token,
	)
	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	response, err :=  c.httpClient.Do(
		req,
	)
	if err != nil{
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
	return nil, fmt.Errorf(
		"github returned status %d",
		response.StatusCode,
	)
}

	defer response.Body.Close()


	var repositoryResponse repositoriesResponse

	err = json.NewDecoder(response.Body).Decode(&repositoryResponse)
	if err != nil{
		return nil, err
	}

	var repositories []Repository

	for _, node := range repositoryResponse.Data.User.Repositories.Nodes {

	repo := Repository{
		Name:        node.Name,
		Description: node.Description,
		Stars:       node.StargazerCount,
		Forks:       node.ForkCount,
		URL:         node.URL,
		CreatedAt:   node.CreatedAt,
		UpdatedAt:   node.UpdatedAt,
	}

	if node.PrimaryLanguage != nil {
		repo.Language = node.PrimaryLanguage.Name
	}

	repositories = append(repositories, repo)
}

return repositories, nil


}


func (c *Client) GetUserContributions(username string) (*Contributions, error){

	reqBody := graphQLRequest{
		Query: getUserContributionsQuery,
		Variables: map[string]any{
			"login" : username,
		},

	}

	body, err := json.Marshal(reqBody)
	if err != nil{
		return nil, err
	}

	req , err := http.NewRequest(
		http.MethodPost,
		"https://api.github.com/graphql",
		bytes.NewBuffer(body),
	)
	if err != nil{
		return nil, err
	}

	req.Header.Set(
		"Authorization",
		"Bearer "+c.token,
	)
	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	response, err := c.httpClient.Do(req)
	if err != nil{
		return nil, err
	}

	defer response.Body.Close()

	var graphqlResponse contributionsResponse

	err = json.NewDecoder(response.Body).Decode(&graphqlResponse)
	if err != nil{
		return nil, err
	}

	contributions := &Contributions{
		TotalCommits: graphqlResponse.Data.User.ContributionsCollection.TotalCommitContributions,
		TotalIssues: graphqlResponse.Data.User.ContributionsCollection.TotalIssueContributions,
		TotalPullRequests: graphqlResponse.Data.User.ContributionsCollection.TotalPullRequestContributions,
		TotalReviews: graphqlResponse.Data.User.ContributionsCollection.TotalPullRequestReviewContributions,
    }

	contributions.TotalContributions = contributions.TotalCommits +
									   contributions.TotalIssues +
									   contributions.TotalPullRequests +
									   contributions.TotalReviews							

	return contributions, nil

}


func (c *Client) GetUserGrowthAnalytics(username string) (*GrowthAnalytics, error){

	reqBody := graphQLRequest{
		Query: getUserGrowthAnalyticsQuery,
		Variables: map[string]any{
			"login": username,
		},
	}

	body, err := json.Marshal(reqBody)
	if err != nil{
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		"https://api.github.com/graphql",
		bytes.NewBuffer(body),
	)
	if err != nil{
		return nil, err
	}

	req.Header.Set(
		"Authorization",
		"Bearer "+ c.token,
	)
	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	response, err := c.httpClient.Do(req)
	if err != nil{
		return nil, err
	}

	defer response.Body.Close()

	var graphqlResponse growthResponse

	err = json.NewDecoder(response.Body).Decode(&graphqlResponse)
	if err != nil{
		return nil, err
	}

	growth := &GrowthAnalytics{
		TotalContributions: graphqlResponse.
							Data.User.
							ContributionsCollection.
							ContributionCalendar.
							TotalContributions,
	}

	activeDays := 0
	bestDays := 0

	for _, week := range graphqlResponse.
						 Data.
						 User.
						 ContributionsCollection.
						 ContributionCalendar.
						 Weeks{

							for _, day := range week.ContributionDays{

								if day.ContributionCount > 0{
									activeDays++
								}

								if day.ContributionCount > bestDays{
									bestDays = day.ContributionCount
								}
							}
						 }

	growth.ActiveDays = activeDays
	growth.BestDayContributions = bestDays
	
	if activeDays > 0 {
		growth.AverageContributionsDay =
			float64(growth.TotalContributions) /
				float64(activeDays)
	}

	return growth, nil
}
