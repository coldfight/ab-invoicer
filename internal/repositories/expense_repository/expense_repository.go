package expense_repository

import (
	"database/sql"
	"github.com/coldfight/ab-invoicer/internal/models"
)

func GetByInvoiceId(invoiceId int, db *sql.DB) (models.ExpenseList, error) {
	stmt, err := db.Prepare(`
SELECT description, unitPrice, quantity
FROM expenses
WHERE invoice = ?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(invoiceId)
	if err != nil {
		return nil, err
	}

	var expenses models.ExpenseList
	for rows.Next() {
		var expense models.Expense
		err = rows.Scan(&expense.Description, &expense.UnitPrice, &expense.Quantity)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}

	return expenses, nil
}
