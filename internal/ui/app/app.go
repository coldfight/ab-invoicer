package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/coldfight/ab-invoicer/internal/ui/menu"
)

func Run() {
	m := menu.Menu{
		Options: []menu.MenuItem{
			{Text: "new check-in", OnPress: func() tea.Msg { return menu.ToggleCasingMsg{} }},
			{Text: "view check-ins", OnPress: func() tea.Msg { return menu.ToggleCasingMsg{} }},
		},
	}
	p := tea.NewProgram(m)
	if err := p.Start(); err != nil {
		panic(err)
	}
}
