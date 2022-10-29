package invoice_list

import "github.com/coldfight/ab-invoicer/internal/models"

// @todo: Rename this package to db instead of models
type invoicesMsg []models.Invoice
type errMsg struct{ err error }

func (e errMsg) Error() string {
	return e.err.Error()
}
