package main_menu

type errMsg struct{ err error }

func (e errMsg) Error() string {
	return e.err.Error()
}
