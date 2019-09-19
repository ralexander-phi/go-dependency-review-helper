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
	_ = filepath.Walk(os.Getenv("GITHUB_WORKSPACE")+"/", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
