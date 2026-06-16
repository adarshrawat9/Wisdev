package github

import (
	"bytes"
	"encoding/json"
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
		"Bearer" + c.token,
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

