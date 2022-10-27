package menu

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

type ToggleCasingMsg struct{}

type Menu struct {
	Options       []MenuItem
	SelectedIndex int
}

type MenuItem struct {
	Text    string
	OnPress func() tea.Msg
}

func (m Menu) Init() tea.Cmd {
	return nil
}

func (m Menu) View() string {
	var options []string
	for i, o := range m.Options {
		pointer := " "
		if i == m.SelectedIndex {
			pointer = ">"
		}
		options = append(options, fmt.Sprintf("%s %s", pointer, o.Text))
	}
	return fmt.Sprintf(
		"%s \n\nPress enter/return to select a list item, arrow keys to move, or Ctrl+C to exit.",
		strings.Join(options, "\n"),
	)
}

func (m Menu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case ToggleCasingMsg:
		return m.toggleSelectedItemCase(), nil
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "down", "right", "up", "left":
			return m.moveCursor(msg.(tea.KeyMsg)), nil
		case "enter", "return":
			return m, m.Options[m.SelectedIndex].OnPress
		}
	}
	return m, nil
}

func (m Menu) moveCursor(msg tea.KeyMsg) Menu {
	switch msg.String() {
	case "up", "left":
		m.SelectedIndex--
	case "down", "right":
		m.SelectedIndex++
	default:
		// do nothing
	}
	optCount := len(m.Options)
	m.SelectedIndex = (m.SelectedIndex + optCount) % optCount
	return m
}

func (m Menu) toggleSelectedItemCase() tea.Model {
	selectedText := m.Options[m.SelectedIndex].Text
	if selectedText == strings.ToUpper(selectedText) {
		m.Options[m.SelectedIndex].Text = strings.ToLower(selectedText)
	} else {
		m.Options[m.SelectedIndex].Text = strings.ToUpper(selectedText)
	}
	return m
}
