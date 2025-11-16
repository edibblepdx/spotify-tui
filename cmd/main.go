package main

import (
	"fmt"
	"log"
	"os"

	"music/internal/spotify"
	"music/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	CLIENT_ID     string
	CLIENT_SECRET string
)

func init() {
	CLIENT_ID = os.Getenv("CLIENT_ID")
	if CLIENT_ID == "" {
		log.Fatal("no client id in CLIENT_ID")
	}
	CLIENT_SECRET = os.Getenv("CLIENT_SECRET")
	if CLIENT_ID == "" {
		log.Fatal("no client secret in CLIENT_SECRET")
	}
}

func main() {
	token, err := spotify.GetToken(CLIENT_ID, CLIENT_SECRET)
	if err != nil {
		log.Fatal(err)
	}

	p := tea.NewProgram(tui.Model{
		Query: os.Args[1],
		Token: token.AccessToken,
	})

	if _, err := p.Run(); err != nil {
		fmt.Printf("Uh oh, there was an error: %v\n", err)
		os.Exit(1)
	}
}
