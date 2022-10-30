package app

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/coldfight/ab-invoicer/internal/ui/pages/main_menu"
	"os"
)

var p *tea.Program

type sessionState int

const (
	mainMenuView sessionState = iota
	invoiceFormView
	editInvoiceFormView
	invoiceItemView
	invoiceListView
)

// AppModel the main application model which holds/controls other models
type AppModel struct {
	state           sessionState
	mainMenu        tea.Model // MainMenuModel
	invoiceForm     tea.Model // InvoiceFormModel
	editInvoiceForm tea.Model // EditInvoiceFormModel
	invoiceItem     tea.Model // InvoiceItemModel
	invoiceList     tea.Model // InvoiceListModel
	windowSize      tea.WindowSizeMsg
}

// Run the entry point of the application
func Run() {
	app := NewApp()

	p := tea.NewProgram(app, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func NewApp() AppModel {
	return AppModel{
		state:    mainMenuView,
		mainMenu: main_menu.New(),
	}
}

func (app AppModel) Init() tea.Cmd {
	return nil
}

// Update handle IO and commands
func (app AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg.(type) {
	case tea.WindowSizeMsg:
		app.windowSize = msg.(tea.WindowSizeMsg)
	}

	switch app.state {
	case mainMenuView:
		app.mainMenu, cmd = app.mainMenu.Update(msg)
	case invoiceListView:
		app.invoiceList, cmd = app.invoiceList.Update(msg)
	}
	cmds = append(cmds, cmd)
	return app, tea.Batch(cmds...)
}

func (app AppModel) View() string {
	switch app.state {
	case invoiceListView:
		return app.invoiceList.View()
	default:
		return app.mainMenu.View()
	}
}
