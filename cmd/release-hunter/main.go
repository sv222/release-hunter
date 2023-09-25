package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

const searchResultLimit = 10
const cliVersion = "v0.1.0"

func main() {
	var (
		token      string
		tokenAlias string
		repoFlag   string
		repoAlias  string
		findFlag   string
		findAlias  string
		keyword    string
		helpFlag   bool
		versionFlag bool
	)

	// Define command-line flags and their descriptions
	flag.StringVar(&token, "token", "", "GitHub personal access token")
	flag.StringVar(&tokenAlias, "t", "", "GitHub personal access token (alias)")
	flag.StringVar(&repoFlag, "repo", "", "GitHub repository in the format user/repo")
	flag.StringVar(&repoAlias, "r", "", "GitHub repository in the format user/repo (alias)")
	flag.StringVar(&findFlag, "find", "", "Search GitHub repositories by keyword")
	flag.StringVar(&findAlias, "f", "", "Search GitHub repositories by keyword (alias)")
	flag.StringVar(&keyword, "keyword", "", "Filter links by keyword")
	flag.StringVar(&keyword, "k", "", "Filter links by keyword (alias)")
	flag.BoolVar(&helpFlag, "help", false, "Show usage and examples")
	flag.BoolVar(&helpFlag, "h", false, "Show usage and examples (alias)")
	flag.BoolVar(&versionFlag, "version", false, "Show the Release Hunter CLI version")
	flag.BoolVar(&versionFlag, "v", false, "Show the Release Hunter CLI version (alias)")

	// Parse command-line flags
	flag.Parse()

	// If the user requested the version, display it and exit
	if versionFlag {
		fmt.Printf("Version: %s\n", cliVersion)
		return
	}

	// Check if both -find or -f and -repo or -r flags are used together
	if (findFlag != "" || findAlias != "") && (repoFlag != "" || repoAlias != "") {
		fmt.Println("Error: -f and -r flags cannot be used together. Use the -k flag to filter results.")
		os.Exit(1)
	}

	// Check if the GITHUB_TOKEN environment variable is set and use it as the default token if available
	envToken := os.Getenv("GITHUB_TOKEN")
	if envToken != "" {
		token = envToken
	}

	// If the user requested help or if there are no other flags specified, print usage information and exit
	if helpFlag || (token == "" && tokenAlias == "" && repoFlag == "" && repoAlias == "" && findFlag == "" && findAlias == "") {
		printUsage()
		return
	}

	// If the user specified the -find or -f flag, perform a GitHub repository search
	if findFlag != "" || findAlias != "" {
		if findAlias != "" {
			findFlag = findAlias
		}

		// Check if a GitHub token is provided
		if token == "" && tokenAlias == "" {
			fmt.Println("GitHub token is required. Use 'export GITHUB_TOKEN=<token>' or '-t' flag.")
			os.Exit(1)
		}

		// Create a GitHub client with the provided token
		client, err := createGitHubClient(token)
		if err != nil {
			log.Fatal(err)
		}

		ctx := context.Background()
		searchOptions := &github.SearchOptions{
			ListOptions: github.ListOptions{PerPage: searchResultLimit},
		}
		// Use the provided keyword for filtering
		query := findFlag + " " + keyword // Modify the query to include the keyword
		repoResults, _, err := client.Search.Repositories(ctx, query, searchOptions)
		if err != nil {
			log.Fatal(err)
		}

		// Print the top search results
		for i, repo := range repoResults.Repositories {
			if i >= searchResultLimit {
				break
			}
			fmt.Printf("%d. %s - %s\n", i+1, repo.GetFullName(), repo.GetDescription())
		}
		return
	}

	// If neither -find nor -f is specified, expect token and repository information
	if findFlag == "" && findAlias == "" && repoFlag == "" && repoAlias == "" {
		fmt.Println("Arguments are required. Use -help or -h for usage and examples.")
		os.Exit(1)
	}

	// Use alias values if provided
	if tokenAlias != "" {
		token = tokenAlias
	}

	if repoAlias != "" {
		repoFlag = repoAlias
	}

	var owner, repo string
	
	// Check if the repository flag contains a slash
	if strings.Contains(repoFlag, "/") {
		// Split the repository flag into owner and repo name
		repoParts := strings.SplitN(repoFlag, "/", 2)
		if len(repoParts) != 2 || repoParts[0] == "" || repoParts[1] == "" {
			fmt.Println("Invalid repo format. Please use the format user/repo.")
			os.Exit(1)
		}
		owner = repoParts[0]
		repo = repoParts[1]
	} else {
		fmt.Println("Invalid repo format. Please use the format user/repo.")
		os.Exit(1)
	}

	// Create a GitHub client with the provided token
	client, err := createGitHubClient(token)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	release, _, err := client.Repositories.GetLatestRelease(ctx, owner, repo)
	if err != nil {
		log.Fatalf("Error getting latest release: %v\n", err)
	}

	if release == nil {
		fmt.Println("No release found for repo.")
		return
	}

	// Convert keyword to lowercase for case-insensitive filtering
	keyword = strings.ToLower(keyword)

	// Iterate through release assets and filter by keyword
	for _, asset := range release.Assets {
		// Convert asset name to lowercase for case-insensitive comparison
		assetName := strings.ToLower(asset.GetName())
		// Check if the keyword is present in the asset name
		if keyword == "" || strings.Contains(assetName, keyword) {
			fmt.Printf("%s\n", asset.GetBrowserDownloadURL())
		}
	}
}

// Function to create a GitHub client with the provided token
func createGitHubClient(token string) (*github.Client, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return client, nil
}

// Function to print usage information
func printUsage() {
	usage := `Usage: rh [options]

Options:
  -token, -t     GitHub personal access token.
  -repo, -r      GitHub repository in "user/repo" format.
  -find, -f      Search GitHub repositories by keyword.
  -keyword, -k   Filter links by keyword.
  -help, -h      Show usage and examples.
  -version, -v   Show the version of the CLI.

Examples:
  1. Provide your token and the exact name of the user/repository:
	rh -t YOUR_GITHUB_TOKEN -r user/repo
  2. Search for repositories containing the keyword 'helm' to find user/repo for previous example:
	rh -t YOUR_GITHUB_TOKEN -f helm
  3. List assets of a specific repository and filter by keyword 'arm':
	rh -t YOUR_GITHUB_TOKEN -r helm/helm -k arm
  4. A more accurate filter for finding the right repository:
	rh -t YOUR_GITHUB_TOKEN -f helm -k manager

Notice: if the environment variable "GITHUB_TOKEN" is set, there is no need to use the "-t" flag.
Visit: github.com/sv222/release-hunter`
	fmt.Println(usage)
}