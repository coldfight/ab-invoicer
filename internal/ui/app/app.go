package app

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/coldfight/ab-invoicer/internal/models"
	"github.com/coldfight/ab-invoicer/internal/tools/logit"
	"github.com/coldfight/ab-invoicer/internal/ui/common"
	"github.com/coldfight/ab-invoicer/internal/ui/pages/edit_invoice_form"
	"github.com/coldfight/ab-invoicer/internal/ui/pages/invoice_item"
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
		app.windowSize = msg.(tea.WindowSizeMsg)
	case common.SwitchToStateMsg:
		s := msg.(common.SwitchToStateMsg).State
		switch s {
		case common.InvoiceListView:
			app.invoiceList = invoice_list.New(app.windowSize)
		case common.InvoiceItemView:
			invoiceNumber, ok := msg.(common.SwitchToStateMsg).Data.(models.InvoiceNumber)
			if !ok {
				logit.Error("type asserting invoiceNumber is incorrect", invoiceNumber)
				return app, func() tea.Msg {
					return common.SwitchToStateMsg{State: common.InvoiceListView}
				}
			} else {
				app.invoiceItem = invoice_item.New(invoiceNumber, app.windowSize)
			}
		case common.EditInvoiceFormView:
			invoiceNumber, ok := msg.(common.SwitchToStateMsg).Data.(models.InvoiceNumber)
			if !ok {
				logit.Error("type asserting invoiceNumber is incorrect", invoiceNumber)
				return app, func() tea.Msg {
					return common.SwitchToStateMsg{State: common.InvoiceListView}
				}
			} else {
				app.editInvoiceForm = edit_invoice_form.New(invoiceNumber, app.windowSize)
			}
		}
		app.state = s
	}

	// Delegate the Update methods to each sub-model
	switch app.state {
	case common.InvoiceListView:
		app.invoiceList, cmd = app.invoiceList.Update(msg)
	case common.InvoiceItemView:
		app.invoiceItem, cmd = app.invoiceItem.Update(msg)
	case common.EditInvoiceFormView:
		app.editInvoiceForm, cmd = app.editInvoiceForm.Update(msg)
	default:
		app.mainMenu, cmd = app.mainMenu.Update(msg)
	}
	cmds = append(cmds, cmd)
	return app, tea.Batch(cmds...)
}

func (app AppModel) View() string {
	switch app.state {
	case common.InvoiceListView:
		return app.invoiceList.View()
	case common.InvoiceItemView:
		return app.invoiceItem.View()
	case common.EditInvoiceFormView:
		return app.editInvoiceForm.View()
	default:
		return app.mainMenu.View()
	}
}
