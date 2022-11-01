package common

import "github.com/charmbracelet/lipgloss"

var (
	AppStyle = lipgloss.NewStyle().Padding(1, 2)

	TitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1)

	StatusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Render

	// DocStyle styling for viewports
	DocStyle = lipgloss.NewStyle().Margin(0, 2)

	// HelpStyle styling for help context menu
	HelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render

	// ErrStyle provides styling for error messages
	ErrStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#bd534b")).Render

	// AlertStyle provides styling for alert messages
	AlertStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("62")).Render
)
