package github

import "sort"

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

func (c *Service) GetUserAnalytics(username string) (*Analytics, error){

	repos, err := c.GetUserRepositories(username)
	if err != nil{
		return nil, err
	}

	analytics := &Analytics{}
	languageStats := make(map[string]*LanguageStats)

	for _, repo := range repos {

		analytics.TotalRepositories++
		analytics.TotalStars += repo.Stars
		analytics.TotalForks += repo.Forks

		if repo.Language != "" {

			if _, exists := languageStats[repo.Language]; !exists {
				languageStats[repo.Language] = &LanguageStats{
					Name: repo.Language,
				}
			}

			languageStats[repo.Language].Repositories++
			languageStats[repo.Language].Stars += repo.Stars
			languageStats[repo.Language].Forks += repo.Forks
		}
	}

	// Convert map -> slice
	for _, stats := range languageStats {
		analytics.Languages = append(
			analytics.Languages,
			*stats,
		)
	}

	// Find most used language
	maxRepos := 0

	for _, lang := range analytics.Languages {
		if lang.Repositories > maxRepos {
			maxRepos = lang.Repositories
			analytics.MostUsedLanguage = lang.Name
		}
	}

	// Sort repositories by stars
	sort.Slice(repos, func(i, j int) bool {
		return repos[i].Stars > repos[j].Stars
	})

	limit := 3
	if len(repos) < limit {
		limit = len(repos)
	}

	for i := 0; i < limit; i++ {
		analytics.TopRepositories = append(
			analytics.TopRepositories,
			TopRepository{
				Name:  repos[i].Name,
				Stars: repos[i].Stars,
				URL:   repos[i].URL,
			},
		)
	}

	return analytics, nil

	

}
