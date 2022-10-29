package invoice_list

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/coldfight/ab-invoicer/internal/models"
	"github.com/coldfight/ab-invoicer/internal/services/invoice_service"
)

// @todo: This does not work... Iwill need to figure out how to initialize the list since this submodel's Init() does not get called
type Model struct {
	invoiceList []models.Invoice
	cursor      int
	err         error
}

func getInvoiceList() tea.Msg {
	invoices := invoice_service.GetInvoices()
	return invoicesMsg(invoices)
}

func getInvoice(id int) tea.Cmd {
	return func() tea.Msg {
		invoice := invoice_service.GetFullInvoiceRecord(id)
		return invoicesMsg{invoice}
	}
}

func New() Model {
	return Model{}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case invoicesMsg:
		m.invoiceList = msg
	case errMsg:
		m.err = msg
		return m, tea.Quit
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.invoiceList)-1 {
				m.cursor++
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	if m.err != nil {
		return fmt.Sprintf("\nWe had some trouble: %v\n\n", m.err)
	}

	// Header
	s := "Which invoice do you want to view?\n\n"

	for i, invoice := range m.invoiceList {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s Invoice #%s - %s\n",
			cursor,
			invoice.InvoiceNumber.Padded(),
			invoice.InvoiceDate.Format("2006-01-02"),
		)
	}

	// Footer
	s += "\nPress q to quit.\n"

	return s
}
