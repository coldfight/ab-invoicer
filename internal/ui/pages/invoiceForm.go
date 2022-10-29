package pages

import (
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

type InvoiceFormModel struct {
}

func NewInvoiceFormPage() InvoiceFormModel {
	return InvoiceFormModel{}
}

func (m InvoiceFormModel) Init() tea.Cmd {
	log.Println(m)
	return nil
}

func (m InvoiceFormModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m InvoiceFormModel) View() string {
	return "New Invoice Form"
}
