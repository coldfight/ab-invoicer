package pages

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

// Model Data

type PageB struct {
	text              string
	name              string
	BackToMainHandler func() (tea.Model, tea.Cmd)
}

func (page PageB) GetName() string {
	return page.name
}

func (page PageB) BackToMain() (tea.Model, tea.Cmd) {
	return page.BackToMainHandler()
}

func NewPageB(text string, backToMainHandler func() (tea.Model, tea.Cmd)) PageB {
	return PageB{text: text, name: "Page B", BackToMainHandler: backToMainHandler}
}

func (page PageB) Init() tea.Cmd {
	return nil
}

func (page PageB) View() string {
	textLen := len(page.text) + len(page.name) + 2
	topAndBottomBar := strings.Repeat("*", textLen+4)
	return fmt.Sprintf("%s\n* %s *\n%s\n\nPress Ctrl+C to exit",
		topAndBottomBar,
		fmt.Sprintf("%s: %s", page.name, page.text),
		topAndBottomBar,
	)
}

func (page PageB) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "ctrl+c":
			return page, tea.Quit
		case "esc":
			return page.BackToMain()
		}
	}
	return page, nil
}
