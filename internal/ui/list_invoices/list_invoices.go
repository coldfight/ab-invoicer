package list_invoices

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/coldfight/ab-invoicer/internal/models"
	"github.com/coldfight/ab-invoicer/internal/services/invoice_service"
	"os"
)

const url = "https://amedeobonanno.com/"

type myModel struct {
	invoiceList []models.Invoice
	cursor      int
	err         error
	//selected    map[int]struct{}
}

type errMsg struct{ err error }
type invoicesMsg []models.Invoice

func (e errMsg) Error() string {
	return e.err.Error()
}

func getList() tea.Msg {
	invoices := invoice_service.GetInvoices()
	return invoicesMsg(invoices)
}

//func checkSomeUrl(myUrl string) tea.Cmd {
//	return func() tea.Msg {
//		c := &http.Client{Timeout: 10 * time.Second}
//		res, err := c.Get(myUrl)
//
//		if err != nil {
//			return errMsg{err}
//		}
//
//		return statusMsg(res.StatusCode)
//	}
//}

func (m myModel) Init() tea.Cmd {
	return getList
}

func (m myModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m myModel) View() string {
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

	s += "\nPress q to quit.\n"

	return s
}

func Run() {
	p := tea.NewProgram(myModel{})
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v\n", err)
		os.Exit(1)
	}
}
