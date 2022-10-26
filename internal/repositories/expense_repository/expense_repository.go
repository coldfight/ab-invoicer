package expense_repository

import (
	"github.com/coldfight/ab-invoicer/internal/models"
	"github.com/coldfight/ab-invoicer/internal/services/db_service"
)

func GetByInvoiceId(invoiceId int) (models.ExpenseList, error) {
	db, err := db_service.GetConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

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
	defer rows.Close()

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
