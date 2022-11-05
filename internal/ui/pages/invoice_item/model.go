package invoice_item

import (
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/coldfight/ab-invoicer/internal/models"
	"github.com/coldfight/ab-invoicer/internal/services/invoice_service"
	"github.com/coldfight/ab-invoicer/internal/tools/logit"
	"github.com/coldfight/ab-invoicer/internal/tools/template_helpers"
	"github.com/coldfight/ab-invoicer/internal/ui/common"
	"strconv"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	ownerInfo     table.Model
	customerInfo  table.Model
	expensesTable table.Model
	labourTable   table.Model
	selectedTable int
	table         table.Model
}

func (m model) Init() tea.Cmd { return nil }

type switchTableMsg struct {
	selectedTable table.Model
}

func (m model) SwitchTableCmd(selectedTable int) tea.Cmd {
	return func() tea.Msg {
		switch selectedTable {
		case 0:
			return switchTableMsg{m.ownerInfo}
		case 1:
			return switchTableMsg{m.customerInfo}
		case 2:
			return switchTableMsg{m.expensesTable}
		default:
			return switchTableMsg{m.labourTable}
		}
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case switchTableMsg:
		m.table.Blur()
		m.table = msg.selectedTable
		m.table.Focus()
	case tea.KeyMsg:
		// @todo: CHange these to use a key binding map
		switch msg.String() {
		case "tab":
			logit.Debug("pressing tab so go to next table")
			m.selectedTable += 1
			if m.selectedTable > 3 {
				m.selectedTable = 0
			}
			return m, m.SwitchTableCmd(m.selectedTable)
		case "shift+tab":
			logit.Debug("pressing tab so go to next table")
			m.selectedTable -= 1
			if m.selectedTable < 0 {
				m.selectedTable = 3
			}
			return m, m.SwitchTableCmd(m.selectedTable)
		case "esc":
			backCmd := func() tea.Msg {
				return common.SwitchToStateMsg{State: common.InvoiceListView}
			}
			return m, backCmd
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			logit.Debug(fmt.Sprintf("Let's go to %s!", m.table.SelectedRow()[1]))
			return m, nil
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return common.AppStyle.Render(
		baseStyle.Render(m.table.View()) + "\n",
	)

}

func getDefaultStyles() table.Styles {
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	return s
}

func createOwnerTable(invoice *models.Invoice) table.Model {
	columns := []table.Column{
		{Title: "Field", Width: 10},
		{Title: "Value", Width: 50},
	}
	rows := []table.Row{
		{"Name", invoice.Owner.Name},
		{"Email", invoice.Owner.Email},
		{"Phone", invoice.Owner.Phone},
		{"Street", invoice.Owner.Street},
		{"City", invoice.Owner.City},
		{"Province", invoice.Owner.Province},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(false),
		table.WithHeight(7),
	)

	s := getDefaultStyles()
	t.SetStyles(s)

	return t
}

func createCustomerTable(invoice *models.Invoice) table.Model {
	columns := []table.Column{
		{Title: "Field", Width: 10},
		{Title: "Value", Width: 50},
	}
	rows := []table.Row{
		{"Name", invoice.Customer.Name},
		{"Phone", invoice.Customer.Phone},
		{"Street", invoice.Customer.Street},
		{"City", invoice.Customer.City},
		{"Province", invoice.Customer.Province},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(false),
		table.WithHeight(7),
	)

	s := getDefaultStyles()
	t.SetStyles(s)

	return t
}

func createExpensesTable(invoice *models.Invoice) table.Model {
	columns := []table.Column{
		{Title: "Quantity", Width: 15},
		{Title: "Description", Width: 50},
		{Title: "UnitPrice", Width: 10},
	}

	var rows []table.Row
	for _, e := range invoice.Expenses {
		rows = append(rows, table.Row{
			strconv.Itoa(e.Quantity),
			e.Description,
			template_helpers.AsCurrency(e.UnitPrice),
		})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := getDefaultStyles()
	t.SetStyles(s)

	return t
}

func createLaboursTable(invoice *models.Invoice) table.Model {
	columns := []table.Column{
		{Title: "Date", Width: 15},
		{Title: "Description", Width: 50},
		{Title: "Amount", Width: 10},
	}

	var rows []table.Row
	for _, l := range invoice.Labours {
		rows = append(rows, table.Row{
			l.Date.Format("Jan 02, 2006"),
			l.Description,
			template_helpers.AsCurrency(l.Amount),
		})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := getDefaultStyles()
	t.SetStyles(s)

	return t
}

func New(invoiceNumber models.InvoiceNumber, windowSizeMsg tea.WindowSizeMsg) model {
	invoice := invoice_service.GetFullInvoiceRecord(invoiceNumber)
	owner := createOwnerTable(&invoice)
	owner.Focus()
	customer := createCustomerTable(&invoice)
	expenses := createExpensesTable(&invoice)
	labours := createLaboursTable(&invoice)

	return model{
		selectedTable: 0,
		ownerInfo:     owner,
		customerInfo:  customer,
		expensesTable: expenses,
		labourTable:   labours,
		table:         owner,
	}
}
