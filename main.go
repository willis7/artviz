package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	graph "github.com/awalterschulze/gographviz"
)

// Repository - artifactory repository JSON
type Repository struct {
	Key         string `json:"key"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type"`
	URL         string `json:"url"`
	PackageType string `json:"packageType"`
}

// RepoConf - artifactory repository configuration JSON
type RepoConf struct {
	Key          string   `json:"key"`
	PackageType  string   `json:"packageType"`
	Description  string   `json:"description"`
	Repositories []string `json:"repositories"`
	Rclass       string   `json:"rclass"`
}

var baseURL string

func init() {
	flag.StringVar(&baseURL, "url", "", "Artifactory base url")
	flag.Parse()
}

func main() {
	repos, err := GetRepositories(baseURL)
	if err != nil {
		panic(err)
	}

	g := NewGraph()

	for _, repo := range repos {
		g.AddNode("G", repo.Key, nil)

		conf, err := GetRepoConf(baseURL, repo)
		if err != nil {
			panic(err)
		}

		// Virtual repos have children which make up the group.
		// We iterate over those and create edges in the graph
		if repo.Type == "VIRTUAL" {
			for _, child := range conf.Repositories {
				g.AddEdge(child, repo.Key, true, nil)
			}
		}
	}
	fmt.Println(g.String())
}

// NewGraph returns a graph which escapes special characters
func NewGraph() *graph.Escape {
	g := graph.NewEscape()
	g.SetName("Artifactory")
	g.SetDir(true)
	return g
}

// GetRepositories returns a list of minimal repository details for all repositories
func GetRepositories(baseURL string) ([]Repository, error) {
	resp, err := http.Get(baseURL + "/api/repositories")
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
func GetRepoConf(baseURL string, repo Repository) (RepoConf, error) {
	resp, err := http.Get(baseURL + "/api/repositories/" + repo.Key)
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
