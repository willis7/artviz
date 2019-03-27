package cmd

import (
	"fmt"
	"os"

	graph "github.com/awalterschulze/gographviz"
	"github.com/spf13/cobra"
	"github.com/willis7/artviz/artifactory"
)

var apiKey = ""

var rootCmd = &cobra.Command{
	Use:   "artviz [URL]",
	Short: "CLI for printing your artifactory repos in dot language",
	Long: `This simple command line tool connects to your artifactory instance using the REST API 
and generates a graphviz representation using the DOT language.`,
	// TODO: put some validation here to ensure the arg is a valid url
	Args:    cobra.MinimumNArgs(1),
	Run:     action,
	Version: "v0.1",
}

func init() {
	rootCmd.Flags().StringVarP(&apiKey, "apiKey", "k", "", "Artifactory API Key")
}

// newGraph returns a graph which escapes special characters
func newGraph() *graph.Escape {
	g := graph.NewEscape()
	g.SetName("Artifactory")
	g.SetDir(true)
	return g
}

func action(cmd *cobra.Command, args []string) {
	baseURL := args[0]
	api := artifactory.New(baseURL, apiKey)
	repos, err := api.GetRepositories()
	if err != nil {
		panic(err)
	}

	g := newGraph()
	for _, repo := range repos {
		g.AddNode("G", repo.Key, nil)

		conf, err := api.GetRepoConf(repo)
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

// Execute is the CLI entrypoint
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
