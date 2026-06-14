package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Client struct{
	httpClient *http.Client
}


func NewClient() *Client{
	return &Client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}


func (c *Client) GetUser(username string) (*User, error){
	
	url := fmt.Sprintf("https://api.github.com/users/%s",username,)

	response, err := c.httpClient.Get(url)
	if err != nil{
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusNotFound{
		return nil, errors.New("github user not foung")
	}

	if response.StatusCode != http.StatusOK{
		return nil, fmt.Errorf("github responded with %d status code", response.StatusCode)
	}

	var user User

	err = json.NewDecoder(response.Body).Decode(&user)
	if err != nil{
		return nil, err
	}

	return &user, nil
}