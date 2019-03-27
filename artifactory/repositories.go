package artifactory

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
