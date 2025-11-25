package tui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
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

type model struct {
	search textinput.Model
	list   list.Model
	focus  uint
	token  string
	err    error
}

const (
	focusSearch uint = iota
	focusList
)

func NewModel(token string) model {
	ti := textinput.New()
	ti.Placeholder = "Search..."
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	d := list.NewDefaultDelegate()
	d.SetHeight(11)

	lt := list.New([]list.Item{}, d, 50, 20)
	lt.DisableQuitKeybindings()
	lt.Title = "Tracks"

	return model{
		search: ti,
		list:   lt,
		token:  token,
	}
}

func (m model) Init() tea.Cmd {
	//return searchTracks(m.query, m.token)
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
		cmd  tea.Cmd
	)

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
			switch msg.String() {
			// exit the program.
			case "ctrl+c":
				return m, tea.Quit

			// focus the list.
			case "esc", "tab":
				m.focus = focusList
				m.search.Blur()

			// perform search.
			case "enter":
				return m, searchTracks(m.search.Value(), m.token)
			}

		case focusList:
			switch msg.String() {

			// exit the program.
			case "ctrl+c":
				return m, tea.Quit

			// focus the search.
			case "esc", "tab":
				m.focus = focusSearch
				m.search.SetValue("")
				m.search.Focus()

			// plays and pause
			case " ":
				spotify.PlayPause()

			// selects a song
			case "l", "right":
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

	m.search, cmd = m.search.Update(msg)
	cmds = append(cmds, cmd)

	m.list, cmd = m.list.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
