package common

type SessionState int

const (
	MainMenuView SessionState = iota
	InvoiceFormView
	EditInvoiceFormView
	InvoiceItemView
	InvoiceListView
)

type SwitchToStateMsg struct {
	State        SessionState
	Data         any
	ConstructNew bool
}
