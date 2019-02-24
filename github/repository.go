package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Repository represents a github repository with limited data.
type Repository struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Owner      User   `json:"owner"`
	Stargazers []User `json:"stargazers"`
}

// I _think_ I'm reading these instructions correctly and I need to limit data..
const (
	repositoryLimit          = 3
	repositoryStargazerLimit = 3
	repositorydepth          = 3
	repositoryURI            = "https://api.github.com/users/%v/repos"
	repositoryStargazerURI   = "https://api.github.com/repos/%v/%v/stargazers"
)

// GetRepositories will return the repositories. Non-recursive.
func GetRepositories(login string) (repositories []Repository, err error) {
	request, err := http.NewRequest("GET", fmt.Sprintf(repositoryURI, login), nil)
	if err != nil {
		return
	}
	if token != "" {
		request.Header.Add("Authorization", fmt.Sprintf("token %v", token))
	}

	// Add a timeout to this nonsense...
	http.DefaultClient.Timeout = 10 * time.Second
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)

	decoder.Decode(&repositories)

	if len(repositories) > repositoryLimit {
		repositories = repositories[:repositoryLimit]
	}

	for ii := range repositories {
		repositories[ii].GetStargazers()
	}

	return
}

// GetStargazers will return the stargazers. Non-recursive.
func (repository *Repository) GetStargazers() (repositories []Repository, err error) {
	request, err := http.NewRequest("GET", fmt.Sprintf(repositoryStargazerURI, repository.Owner.Login, repository.Name), nil)
	if err != nil {
		return
	}
	if token != "" {
		request.Header.Add("Authorization", fmt.Sprintf("token %v", token))
	}

	// Add a timeout to this nonsense...
	http.DefaultClient.Timeout = 10 * time.Second
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)

	decoder.Decode(&repository.Stargazers)

	if len(repository.Stargazers) > repositoryStargazerLimit {
		repository.Stargazers = repository.Stargazers[:repositoryStargazerLimit]
	}

	return
}

// GetRepositoriesRecursive is written for this challenge.
func GetRepositoriesRecursive(login string) (user User, err error) {
	user.Login = login

	user.Repositories, err = recurseGetRepositories(user, 0)

	return
}

func recurseGetRepositories(user User, depth int) (repositories []Repository, err error) {
	if depth >= repositorydepth {
		return nil, err
	}

	repositories, err = GetRepositories(user.Login)
	if err != nil {
		return
	}

	for ii := range repositories {
		for iii := range repositories[ii].Stargazers {
			repositories[ii].Stargazers[iii].Repositories, err = recurseGetRepositories(repositories[ii].Stargazers[iii], depth+1)
			if err != nil {
				return
			}
		}
	}

	return
}
