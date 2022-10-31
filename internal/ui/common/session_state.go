package common

type SessionState int

const (
	MainMenuView SessionState = iota
	InvoiceFormView
	EditInvoiceFormView
	InvoiceItemView
	InvoiceListView
)

type SwitchToViewMsg struct{ View SessionState }
