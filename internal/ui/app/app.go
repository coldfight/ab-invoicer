package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/coldfight/ab-invoicer/internal/ui/pages"
	"log"
)

type AppModel struct {
	state       pages.State
	mainMenu    pages.MainMenuModel
	invoiceForm pages.InvoiceFormModel
}

func (app AppModel) Init() tea.Cmd {
	return nil
}

func (app AppModel) SetPage(s pages.State) (tea.Model, tea.Cmd) {
	app.state = s
	cmd := func() tea.Msg {
		return nil
	}
	return app, cmd
}

func (app AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "ctrl+c":
			log.Println("Apps quit")
			return app, tea.Quit
		}
	}

	switch app.state {
	case pages.ShowMainMenu:
		log.Println("show main menu message")
		m, cmd := app.mainMenu.Update(msg)
		mainMenu, ok := m.(pages.MainMenuModel)
		if ok {
			app.mainMenu = mainMenu
			// @todo: Logic for changing state goes here. You change state based on how the update affected the menu model
			return app, cmd
		}
	case pages.ShowInvoiceForm:
		log.Println("show invoice form message")
		m, cmd := app.invoiceForm.Update(msg)
		invoiceForm, ok := m.(pages.InvoiceFormModel)
		if ok {
			app.invoiceForm = invoiceForm
			// @todo: Logic for changing state goes here. You change state based on how the update affected the invoice form model
			return app, cmd
		}
	}

	return app, nil
}

func (app AppModel) View() string {
	switch app.state {
	case pages.ShowInvoiceForm:
		return app.invoiceForm.View()
	default:
		return app.mainMenu.View()
	}
}

func NewApp() AppModel {
	return AppModel{
		state:       pages.ShowMainMenu,
		mainMenu:    pages.NewMainMenuPage(),
		invoiceForm: pages.NewInvoiceFormPage(),
	}
}

func Run() {
	app := NewApp()

	p := tea.NewProgram(app, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		panic(err)
	}
}
