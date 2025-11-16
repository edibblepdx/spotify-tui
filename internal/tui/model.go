package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"

	"music/internal/spotify"
)

func searchArtists(query, token string) tea.Cmd {
	return func() tea.Msg {
		items, err := spotify.SearchArtists(query, token)
		if err != nil {
			return errMsg{err}
		}
		return searchArtistMsg(items)
	}
}

func searchTracks(query, token string) tea.Cmd {
	return func() tea.Msg {
		items, err := spotify.SearchTracks(query, token)
		if err != nil {
			return errMsg{err}
		}
		return searchTrackMsg(items)
	}
}

type searchArtistMsg []spotify.ArtistObject
type searchTrackMsg []spotify.TrackObject
type errMsg struct{ err error }

func (e errMsg) Error() string { return e.err.Error() }

type Model struct {
	Query   string
	Token   string
	choices []spotify.TrackObject
	cursor  int
	err     error
}

func (m Model) Init() tea.Cmd {
	//return searchArtists(m.Query, m.Token)
	return searchTracks(m.Query, m.Token)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case errMsg:
		m.err = msg
		return m, tea.Quit

	/*
		case searchArtistMsg:
			for _, item := range msg {
				m.choices = append(m.choices, item.Name)
			}
	*/

	case searchTrackMsg:
		for _, item := range msg {
			m.choices = append(m.choices, item)
		}

	case tea.KeyMsg:

		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The space key plays and pauses
		case " ":
			spotify.PlayPause()

		// The enter key selects a song
		case "enter":
			spotify.OpenUri(m.choices[m.cursor].Uri)

		}
	}

	return m, nil
}

func (m Model) View() string {
	// Header
	s := "Artists\n\n"

	for i, choice := range m.choices {

		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		s += fmt.Sprintf("%s %s ~ %s\n",
			cursor, choice.Name, choice.Artists[0].Name)
	}

	// Footer
	s += "\nPress q to quit, space to pause, enter to select.\n"

	return s
}
