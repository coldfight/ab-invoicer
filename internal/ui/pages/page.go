package pages

import tea "github.com/charmbracelet/bubbletea"

type Page interface {
	GetName() string
	BackToMain() (tea.Model, tea.Cmd)
	tea.Model
}
