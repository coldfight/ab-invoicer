package invoice_list

import "github.com/coldfight/ab-invoicer/internal/models"

type invoicesMsg []models.Invoice
type errMsg struct{ err error }

func (e errMsg) Error() string {
	return e.err.Error()
}
