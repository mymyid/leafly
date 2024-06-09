package ghupload

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/google/go-github/v59/github"

	"golang.org/x/oauth2"
)

func GithubUpload(ghcreds GHCreds, base64Content string, githubOrg string, githubRepo string, pathFile string, replace bool) (content *github.RepositoryContentResponse, response *github.Response, err error) {
	// Decode the base64 string to byte slice
	fileContent, err := base64.StdEncoding.DecodeString(base64Content)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode base64 content: %w", err)
	}

	// Konfigurasi koneksi ke GitHub menggunakan token akses
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ghcreds.GitHubAccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// Membuat opsi untuk mengunggah file
	opts := &github.RepositoryContentFileOptions{
		Message: github.String("Upload file"),
		Content: fileContent,
		Branch:  github.String("main"),
		Author: &github.CommitAuthor{
			Name:  github.String(ghcreds.GitHubAuthorName),
			Email: github.String(ghcreds.GitHubAuthorEmail),
		},
	}

	// Membuat permintaan untuk mengunggah file
	content, response, err = client.Repositories.CreateFile(ctx, githubOrg, githubRepo, pathFile, opts)
	if (err != nil) && (replace) {
		currentContent, _, _, _ := client.Repositories.GetContents(ctx, githubOrg, githubRepo, pathFile, nil)
		opts.SHA = github.String(currentContent.GetSHA())
		content, response, err = client.Repositories.UpdateFile(ctx, githubOrg, githubRepo, pathFile, opts)
		return
	}

	return
}
