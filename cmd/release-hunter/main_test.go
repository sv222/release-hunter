package main

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/google/go-github/github"
)

// Define a mock struct that satisfies the *github.Client interface
type MockGitHubClient struct{}

func (m *MockGitHubClient) Repositories() *MockRepositoriesService {
	return &MockRepositoriesService{}
}

func (m *MockGitHubClient) Search() *MockSearchService {
	return &MockSearchService{}
}

type MockRepositoriesService struct{}

func (m *MockRepositoriesService) GetLatestRelease(ctx context.Context, owner, repo string) (*github.RepositoryRelease, *github.Response, error) {
	// Implement the mock logic here for GetLatestRelease
	return nil, nil, nil
}

type MockSearchService struct{}

func (m *MockSearchService) Repositories(ctx context.Context, query string, opts *github.SearchOptions) (*github.RepositoriesSearchResult, *github.Response, error) {
	// Implement the mock logic here for Repositories search
	return nil, nil, nil
}

// Implement the createMockGitHubClient function to return an error for an empty token
func createMockGitHubClient(token string) (*MockGitHubClient, error) {
	if token == "" {
		return nil, errors.New("empty token")
	}
	return &MockGitHubClient{}, nil
}

// Define a mock implementation for parseRepositoryFlag
func parseRepositoryFlag(repoFlag string) (string, string, error) {
	parts := strings.Split(repoFlag, "/")
	if len(parts) != 2 {
		return "", "", errors.New("invalid repo flag")
	}
	return parts[0], parts[1], nil
}

// Define a mock implementation for filterByKeyword
func filterByKeyword(assets []github.ReleaseAsset, keyword string) []github.ReleaseAsset {
	var filtered []github.ReleaseAsset
	for _, asset := range assets {
		if strings.Contains(*asset.Name, keyword) {
			filtered = append(filtered, asset)
		}
	}
	return filtered
}

func TestCreateGitHubClient(t *testing.T) {
	tests := []struct {
		name  string
		token string
	}{
		{
			name:  "ValidToken",
			token: "validToken",
			},
			{
			name:  "EmptyToken",
			token: "",
			},
			}

			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					client, err := createMockGitHubClient(tt.token)
					if tt.token != "" {
						if err != nil {
							t.Errorf("createMockGitHubClient(%s) error = %v; want nil", tt.token, err)
						}
						if client == nil {
							t.Error("createMockGitHubClient() returned a nil client; want non-nil")
						}
					} else {
						if err == nil {
							t.Error("createMockGitHubClient() error = nil; want non-nil error")
						}
						if client != nil {
							t.Error("createMockGitHubClient() returned a non-nil client; want nil")
						}
					}
				})
			}
}

func TestParseRepositoryFlag(t *testing.T) {
	tests := []struct {
		name         string
		repoFlag     string
		expectedErr  bool
		expectedUser string
		expectedRepo string
	}{
		{
			name:         "ValidRepoFlag",
			repoFlag:     "user/repo",
			expectedErr:  false,
			expectedUser: "user",
			expectedRepo: "repo",
			},
			{
			name:         "EmptyRepoFlag",
			repoFlag:     "",
			expectedErr:  true,
			expectedUser: "",
			expectedRepo: "",
			},
			{
			name:         "InvalidRepoFlag",
			repoFlag:     "invalid_format",
			expectedErr:  true,
			expectedUser: "",
			expectedRepo: "",
			},
			}

			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					user, repo, err := parseRepositoryFlag(tt.repoFlag)
					if (err != nil) != tt.expectedErr {
						t.Errorf("parseRepositoryFlag(%s) error = %v; want error %v", tt.repoFlag, err, tt.expectedErr)
					}
					if user != tt.expectedUser {
						t.Errorf("parseRepositoryFlag(%s) user = %s; want %s", tt.repoFlag, user, tt.expectedUser)
					}
					if repo != tt.expectedRepo {
						t.Errorf("parseRepositoryFlag(%s) repo = %s; want %s", tt.repoFlag, repo, tt.expectedRepo)
					}
				})
			}
}

func TestFilterByKeyword(t *testing.T) {
	tests := []struct {
		name           string
		assets         []github.ReleaseAsset
		keyword        string
		expectedAssets []github.ReleaseAsset
	}{
		{
			name: "KeywordMatch",
			assets: []github.ReleaseAsset{
				{ID: github.Int64(1), Name: github.String("asset1")},
				{ID: github.Int64(2), Name: github.String("keywordMatch")},
				{ID: github.Int64(3), Name: github.String("asset3")},
				},
				keyword: "keyword",
				expectedAssets: []github.ReleaseAsset{
				{ID: github.Int64(2), Name: github.String("keywordMatch")},
				},
				},
				{
			name: "NoKeywordMatch",
			assets: []github.ReleaseAsset{
				{ID: github.Int64(1), Name: github.String("asset1")},
				{ID: github.Int64(2), Name: github.String("asset2")},
				{ID: github.Int64(3), Name: github.String("asset3")},
				},
				keyword:        "nonexistent",
				expectedAssets: nil,
				},
				}

				for _, tt := range tests {
					t.Run(tt.name, func(t *testing.T) {
						filteredAssets := filterByKeyword(tt.assets, tt.keyword)
						if len(filteredAssets) != len(tt.expectedAssets) {
							t.Errorf("filterByKeyword() returned %d assets; want %d", len(filteredAssets), len(tt.expectedAssets))
						}
						for i, asset := range filteredAssets {
							if *asset.ID != *tt.expectedAssets[i].ID || *asset.Name != *tt.expectedAssets[i].Name {
								t.Errorf("filterByKeyword() asset = %+v; want %+v", asset, tt.expectedAssets[i])
							}
						}
					})
				}
}