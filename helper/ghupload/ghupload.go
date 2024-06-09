package ghupload

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/google/go-github/v59/github"

	"golang.org/x/oauth2"
)

func GithubUploadJPG(ghcreds GHCreds, base64Content string, githubOrg string, githubRepo string, pathFile string, replace bool) (content *github.RepositoryContentResponse, response *github.Response, fileHash string, err error) {
	// Decode the base64 string to byte slice
	fileContent, err := base64.StdEncoding.DecodeString(base64Content)
	if err != nil {
		return nil, nil, "", fmt.Errorf("failed to decode base64 content: %w", err)
	}

	// Calculate hash of the file content +"/#HASHFILE#.jpg"
	fileHash = CalculateHash(fileContent)
	pathFile = pathFile + "/" + fileHash + ".jpg"

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

// Function to calculate the SHA-256 hash of an image
func CalculateHash(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}
