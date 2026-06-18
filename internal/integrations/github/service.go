package github


import(

)

type Service struct{
	client *Client
}


func NewService(client *Client) *Service{
	return &Service{
		client: client,
	}
}


func (s *Service) GetUser(username string) (*User, error) {
	return s.client.GetUser(username)
}

func (s *Service) GetUserRepositories(username string) ([]Repository, error){
	return s.client.GetUserRepositories(username)
}