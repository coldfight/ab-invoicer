package pages

import tea "github.com/charmbracelet/bubbletea"

type Page interface {
	GetName() string
	BackToMain() (tea.Model, tea.Cmd)
	tea.Model
}

type PageId int

const (
	MainMenuPageId PageId = iota
	NewInvoiceFormId
	InvoiceListPageId
	InvoiceItemPageId
)

type State int

const (
	ShowMainMenu State = iota
	ShowInvoiceForm
)
