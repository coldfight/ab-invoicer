package pages

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"strings"
)

type switchPageMsg struct {
	page State
}

type MenuItem struct {
	Text    string
	OnPress func() tea.Msg
}

type MainMenuModel struct {
	Options       []MenuItem
	SelectedIndex int
}

func NewMainMenuPage() MainMenuModel {
	return MainMenuModel{
		Options: []MenuItem{
			{Text: "Create new invoice", OnPress: func() tea.Msg { return switchPageMsg{page: ShowInvoiceForm} }},
			{Text: "View all invoices", OnPress: func() tea.Msg { return switchPageMsg{page: ShowInvoiceForm} }},
		},
	}
}

func (m MainMenuModel) Init() tea.Cmd {
	log.Println(m)
	return nil
}

func (m MainMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Println(m.Options)
	switch msg.(type) {
	case switchPageMsg:
		return m.toggleSelectedItemCase(msg.(switchPageMsg).page), nil
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

func (m MainMenuModel) View() string {
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

func (m MainMenuModel) moveCursor(msg tea.KeyMsg) MainMenuModel {
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

func (m MainMenuModel) toggleSelectedItemCase(page State) tea.Model {
	log.Println("I really need to switch to the selected page at this point...")
	selectedText := m.Options[m.SelectedIndex].Text
	if selectedText == strings.ToUpper(selectedText) {
		m.Options[m.SelectedIndex].Text = strings.ToLower(selectedText)
	} else {
		m.Options[m.SelectedIndex].Text = strings.ToUpper(selectedText)
	}
	return m
}
