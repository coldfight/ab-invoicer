package app

import "github.com/coldfight/ab-invoicer/internal/ui/views"

func Run() {
	//m := menu.Menu{
	//	Options: []menu.MenuItem{
	//		{Text: "Create new invoice", OnPress: func() tea.Msg { return menu.ToggleCasingMsg{} }},
	//		{Text: "View all invoices", OnPress: func() tea.Msg { return menu.ToggleCasingMsg{} }},
	//	},
	//}
	//p := tea.NewProgram(m, tea.WithAltScreen())
	//if err := p.Start(); err != nil {
	//	panic(err)
	//}

	views.Run()
}
