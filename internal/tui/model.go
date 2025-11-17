package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"music/internal/spotify"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

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
	search textinput.Model
	list   list.Model
	focus  int
	token  string
	err    error
}

const (
	focusSearch int = iota
	focusList
)

func InitialModel(token string) Model {
	ti := textinput.New()
	ti.Placeholder = "Search..."
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	lt := list.New([]list.Item{}, list.NewDefaultDelegate(), 50, 20)
	lt.DisableQuitKeybindings()
	lt.Title = "Tracks"

	return Model{
		search: ti,
		list:   lt,
		token:  token,
	}
}

func (m Model) Init() tea.Cmd {
	//return searchTracks(m.query, m.token)
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case errMsg:
		m.err = msg
		return m, tea.Quit

	case searchTrackMsg:
		items := make([]list.Item, len(msg))
		for i, v := range msg {
			items[i] = v
		}
		m.list.SetItems(items)
		m.focus = focusList
		m.search.Blur()

	case tea.KeyMsg:
		switch m.focus {

		case focusSearch:
			switch msg.Type {
			// exit the program.
			case tea.KeyCtrlC:
				return m, tea.Quit

			// focus the list.
			case tea.KeyTab:
				m.focus = focusList
				m.search.Blur()

			// perform search.
			case tea.KeyEnter:
				return m, searchTracks(m.search.Value(), m.token)
			}

		case focusList:
			switch msg.Type {

			// exit the program.
			case tea.KeyCtrlC:
				return m, tea.Quit

			// focus the search.
			case tea.KeyTab:
				m.focus = focusSearch
				m.search.Focus()

			// plays and pause
			case tea.KeySpace:
				spotify.PlayPause()

			// selects a song
			case tea.KeyEnter:
				if track, ok :=
					m.list.SelectedItem().(spotify.TrackObject); ok {
					spotify.OpenUri(track.Uri)
				}
			}
		}

		/*
			case tea.WindowSizeMsg:
				h, v := docStyle.GetFrameSize()
				m.list.SetSize(msg.Width-h, msg.Height-v)

		*/
	}

	var cmd1, cmd2 tea.Cmd

	m.search, cmd1 = m.search.Update(msg)
	m.list, cmd2 = m.list.Update(msg)

	return m, tea.Batch(cmd1, cmd2)
}

func (m Model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s",
		docStyle.Render(m.search.View()),
		docStyle.Render(m.list.View()),
	)
	/*
		return fmt.Sprintf(
			"%s\n\n%s",
			m.search.View(),
			m.list.View(),
		)
	*/
}
