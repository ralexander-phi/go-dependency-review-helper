package main

import (
	//"context"
	"fmt"
	//"github.com/google/go-github/github"
	//"golang.org/x/oauth2"
	"bufio"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Started!")

	fmt.Println(os.Getenv("HOME"))
	fmt.Println(os.Getenv("GITHUB_REF"))
	fmt.Println(os.Getenv("GITHUB_SHA"))
	fmt.Println(os.Getenv("GITHUB_REPOSITORY"))
	fmt.Println(os.Getenv("GITHUB_ACTOR"))
	fmt.Println(os.Getenv("GITHUB_WORKFLOW"))
	fmt.Println(os.Getenv("GITHUB_HEAD_REF"))
	fmt.Println(os.Getenv("GITHUB_BASE_REF"))
	fmt.Println(os.Getenv("GITHUB_EVENT_NAME"))
	fmt.Println(os.Getenv("GITHUB_WORKSPACE"))
	fmt.Println(os.Getenv("GITHUB_ACTION"))
	fmt.Println(os.Getenv("GITHUB_EVENT_PATH"))
	fmt.Println(os.Getenv("RUNNER_OS"))
	fmt.Println(os.Getenv("RUNNER_TOOL_CACHE"))
	fmt.Println(os.Getenv("RUNNER_TEMP"))
	fmt.Println(os.Getenv("RUNNER_WORKSPACE"))

	fmt.Println("")
	fmt.Println("Event Path contents")
	file, _ := os.Open(os.Getenv("GITHUB_EVENT_PATH"))
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())

	}

	fmt.Println("")
	fmt.Println("WORKSPACE")
	var files []string
	err := filepath.Walk(os.Getenv("GITHUB_WORKSPACE")+"/", func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}

	/*
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
	*/
}
