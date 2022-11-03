package invoice_item

type errMsg struct{ err error }

func (e errMsg) Error() string {
	return e.err.Error()
}
