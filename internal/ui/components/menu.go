package components

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"strings"
)

type toggleCaseMsg struct{}

type MenuItem struct {
	Text    string
	OnPress func() tea.Msg
}

type MenuModel struct {
	Options       []MenuItem
	SelectedIndex int
}

func NewMainMenuPage() MenuModel {
	return MenuModel{
		Options: []MenuItem{
			{Text: "Create new invoice", OnPress: func() tea.Msg { return toggleCaseMsg{} }},
			{Text: "View all invoices", OnPress: func() tea.Msg { return toggleCaseMsg{} }},
		},
	}
}

func (m MenuModel) Init() tea.Cmd {
	log.Println(m)
	return nil
}

func (m MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Println(m.Options)
	switch msg.(type) {
	case toggleCaseMsg:
		return m.toggleSelectedItemCase(), nil
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "ctrl+c":
			return m, tea.Quit
		case "down", "right", "up", "left":
			return m.moveCursor(msg.(tea.KeyMsg)), nil
		case "enter", "return":
			return m, m.Options[m.SelectedIndex].OnPress
		}
	}
	return m, nil
}

func (m MenuModel) View() string {
	log.Println(m.Options)

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

func (m MenuModel) moveCursor(msg tea.KeyMsg) MenuModel {
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

func (m MenuModel) toggleSelectedItemCase() tea.Model {
	selectedText := m.Options[m.SelectedIndex].Text
	if selectedText == strings.ToUpper(selectedText) {
		m.Options[m.SelectedIndex].Text = strings.ToLower(selectedText)
	} else {
		m.Options[m.SelectedIndex].Text = strings.ToUpper(selectedText)
	}
	return m
}
