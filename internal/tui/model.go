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

type searchArtistMsg []spotify.ArtistObject
type errMsg struct{ err error }

func (e errMsg) Error() string { return e.err.Error() }

type Model struct {
	Query   string
	Token   string
	choices []string
	cursor  int
	err     error
}

func (m Model) Init() tea.Cmd {
	return searchArtists(m.Query, m.Token)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case errMsg:
		m.err = msg
		return m, tea.Quit

	case searchArtistMsg:
		for _, item := range msg {
			m.choices = append(m.choices, item.Name)
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

		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	// Footer
	s += "\nPress q to quit.\n"

	return s
}
