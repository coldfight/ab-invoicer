package app

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/coldfight/ab-invoicer/internal/tools/logit"
	"github.com/coldfight/ab-invoicer/internal/ui/common"
	"github.com/coldfight/ab-invoicer/internal/ui/pages/invoice_list"
	"github.com/coldfight/ab-invoicer/internal/ui/pages/main_menu"
	"os"
)

var p *tea.Program

// AppModel the main application model which holds/controls other models
type AppModel struct {
	state           common.SessionState
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
		state:    common.MainMenuView,
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

	// Listen for messages
	switch msg.(type) {
	case tea.WindowSizeMsg:
		logit.Debug("Resizing window", msg)
		app.windowSize = msg.(tea.WindowSizeMsg)
	case common.SwitchToStateMsg:
		logit.Debug("Switching State", msg)
		v := msg.(common.SwitchToStateMsg).State
		switch v {
		case common.InvoiceListView:
			logit.Debug("Creating a new invoice list view", msg)
			app.invoiceList = invoice_list.New()
		}
		app.state = v
	}

	// Delegate the Update methods to each sub-model
	switch app.state {
	case common.MainMenuView:
		logit.Debug("Update main menu", msg)
		app.mainMenu, cmd = app.mainMenu.Update(msg)
	case common.InvoiceListView:
		logit.Debug("Update invoice list", msg)
		app.invoiceList, cmd = app.invoiceList.Update(msg)
	default:
		app.mainMenu, cmd = app.mainMenu.Update(msg)
		logit.Warn(fmt.Sprintf("This state - %d - is not yet handled", app.state))
	}
	cmds = append(cmds, cmd)
	return app, tea.Batch(cmds...)
}

func (app AppModel) View() string {
	switch app.state {
	case common.InvoiceListView:
		return app.invoiceList.View()
	default:
		return app.mainMenu.View()
	}
}
