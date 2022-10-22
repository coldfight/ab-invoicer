package receipt_creator

import (
	"github.com/coldfight/ab-invoicer/internal/tools"
)

func Create() {
	templateData := struct {
		Expenses   []string
		GetAbsPath func(string) string
	}{
		Expenses:   []string{"Lysol Aerosol", "Windex Window Cleaner", "Toilet Bowl Cleaner"},
		GetAbsPath: tools.FullFilePath,
	}

	tools.CreatePdf("./templates/test.tmpl", "./storage/pdf.pdf", templateData)
}
