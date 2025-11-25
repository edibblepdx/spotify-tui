package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	docStyle = lipgloss.NewStyle().
			Margin(1, 2)

	appNameStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("2")).
			Foreground(lipgloss.Color("233")).
			Margin(0, 2).
			Padding(0, 1)
)

func (m model) View() string {
	/*
		return fmt.Sprintf(
			"%s\n\n%s",
			docStyle.Render(m.search.View()),
			docStyle.Render(m.list.View()),
		)
	*/
	return docStyle.Render(fmt.Sprintf(
		"%s\n\n%s\n\n%s",
		appNameStyle.Render("Title"),
		m.search.View(),
		m.list.View(),
	))
}
