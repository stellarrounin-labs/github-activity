package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Event struct {
	ID      string  `json:"id"`
	Type    string  `json:"type"`
	Repo    Repo    `json:"repo"`
	Payload Payload `json:"payload"`
}

type Repo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Payload struct {
	Action      string `json:"action"`
	Description string `json:"description"`
}

func main() {
	username := os.Args[1]

	if username == "" {
		fmt.Print("Username not provided")
		return
	}

	requestUrl := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	client := &http.Client{}

	resp, err := client.Get(requestUrl)
	if err != nil {
		fmt.Printf("=== An error has occurred ===\n%s", err)
		return
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	var userEvents []Event
	err = json.NewDecoder(resp.Body).Decode(&userEvents)
	if err != nil {
		fmt.Printf("=== ERROR ===\nUser %s not found", username)
		return
	}

	for _, event := range userEvents {
		fmt.Print("- ")

		switch event.Type {
		case "IssuesEvent":
			fmt.Printf("%s an issue in repo %s\n", event.Payload.Action, event.Repo.Name)
		case "WatchEvent":
			fmt.Printf("starred repo %s\n", event.Repo.Name)
		case "IssueCommentEvent":
			fmt.Printf("%s comment issue in repo %s\n", event.Payload.Action, event.Repo.Name)
		case "CreateEvent":
			fmt.Printf("created repo %s: %s\n", event.Repo.Name, event.Payload.Description)
		case "PushEvent":
			fmt.Printf("pushed to repo %s\n", event.Repo.Name)
		case "PullRequestEvent":
			fmt.Printf("%s pull request in repo %s\n", event.Payload.Action, event.Repo.Name)
		default:
			fmt.Printf("%s action performed in repo %s\n", event.Payload.Action, event.Repo.Name)
		}
	}
}
