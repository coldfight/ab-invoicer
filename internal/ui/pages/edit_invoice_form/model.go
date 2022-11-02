package edit_invoice_form

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/coldfight/ab-invoicer/internal/models"
)

type Model struct {
}

func New(invoiceNumber models.InvoiceNumber, windowSize tea.WindowSizeMsg) Model {
	return Model{}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	return "Edit Invoice Form"
}
