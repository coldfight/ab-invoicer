package pages

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

// Model Data

type PageC struct {
	text              string
	name              string
	BackToMainHandler func() (tea.Model, tea.Cmd)
}

func (page PageC) GetName() string {
	return page.name
}

func (page PageC) BackToMain() (tea.Model, tea.Cmd) {
	return page.BackToMainHandler()
}

func NewPageC(text string, backToMainHandler func() (tea.Model, tea.Cmd)) PageC {
	return PageC{text: text, name: "Page C", BackToMainHandler: backToMainHandler}
}

func (page PageC) Init() tea.Cmd {
	return nil
}

func (page PageC) View() string {
	textLen := len(page.text) + len(page.name) + 2
	topAndBottomBar := strings.Repeat("*", textLen+4)
	return fmt.Sprintf("%s\n* %s *\n%s\n\nPress Ctrl+C to exit",
		topAndBottomBar,
		fmt.Sprintf("%s: %s", page.name, page.text),
		topAndBottomBar,
	)
}

func (page PageC) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
