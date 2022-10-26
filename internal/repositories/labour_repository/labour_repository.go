package labour_repository

import (
	"database/sql"
	"github.com/coldfight/ab-invoicer/internal/models"
	"log"
)

func GetByInvoiceId(invoiceId int, db *sql.DB) (models.LabourList, error) {
	stmt, err := db.Prepare(`
SELECT description, amount, date
FROM labour 
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

	var labourList models.LabourList
	for rows.Next() {
		var labour models.Labour
		var dateStr string
		err = rows.Scan(&labour.Description, &labour.Amount, &dateStr)
		labour.Date.SetFromString("2006-01-02", dateStr)

		if err != nil {
			log.Fatal(err)
		}
		labourList = append(labourList, labour)
	}

	return labourList, nil
}
