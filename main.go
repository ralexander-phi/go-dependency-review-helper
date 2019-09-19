package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"os"
)

func main() {
	fmt.Println("Started!")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)
	fmt.Println("Client created!")

	client := github.NewClient(tc)
	fmt.Println("Github client created!")

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List(ctx, "", nil)

	fmt.Println(repos)
	fmt.Println(err)
}
