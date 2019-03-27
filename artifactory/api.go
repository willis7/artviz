package artifactory

import (
	"encoding/json"
	"net/http"
)

type Client struct {
	client  http.Client
	baseURL string
	apiKey  string
}

func (c Client) get(resource string) (*http.Request, error) {
	req, err := http.NewRequest("GET", c.baseURL+resource, nil)
	if c.apiKey != "" {
		req.Header.Add("X-JFrog-Art-Api", c.apiKey)
	}
	return req, err
}

// GetRepositories returns a list of minimal repository details for all repositories
func (c Client) GetRepositories() ([]Repository, error) {
	req, err := c.get("/api/repositories")
	resp, err := c.client.Do(req)
	if err != nil {
		return []Repository{}, err
	}
	defer resp.Body.Close()

	var repos []Repository
	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		return []Repository{}, err
	}
	return repos, nil
}

// GetRepoConf retrieves partial configuration for a repository
func (c Client) GetRepoConf(repo Repository) (RepoConf, error) {
	req, err := c.get("/api/repositories/" + repo.Key)
	resp, err := c.client.Do(req)
	if err != nil {
		return RepoConf{}, err
	}
	defer resp.Body.Close()

	var conf RepoConf
	err = json.NewDecoder(resp.Body).Decode(&conf)
	if err != nil {
		return RepoConf{}, err
	}
	return conf, nil
}

// New returns a new Artifactory Client
func New(baseURL, apiKey string) Client {
	return Client{
		baseURL: baseURL,
		apiKey:  apiKey,
		client:  http.Client{},
	}
}
