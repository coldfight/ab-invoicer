package app

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/coldfight/ab-invoicer/internal/ui/pages"
	"os"
)

type myApp struct {
}

func (app myApp) Init() tea.Cmd {
	return nil
}

func (app myApp) View() string {
	return "My App..."
}

func (app myApp) backToMainHandler() (tea.Model, tea.Cmd) {
	return app, nil
}

func (app myApp) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "ctrl+c":
			return app, tea.Quit
		case "A":
			page := pages.NewPageA("I'm a new page A yo...", app.backToMainHandler)
			return page, nil
		case "B":
			page := pages.NewPageB("I'm a new page B yo...", app.backToMainHandler)
			return page, nil
		case "C":
			page := pages.NewPageC("I'm a new page C yo...", app.backToMainHandler)
			return page, nil
		}
	}

	return app, nil
}

func Run() {
	if f, err := tea.LogToFile("logs/debug.log", "app"); err != nil {
		fmt.Println("Couldn't open a file for logging:", err)
		os.Exit(1)
	} else {
		defer f.Close()
	}

	p := tea.NewProgram(myApp{}, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		panic(err)
	}
}
