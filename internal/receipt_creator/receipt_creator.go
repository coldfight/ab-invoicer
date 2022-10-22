package receipt_creator

import (
	"github.com/coldfight/ab-invoicer/internal/tools"
	"html/template"
)

func Create() {
	templateData := struct {
		Data                string
		Expenses            []string
		BootstrapStylesheet template.CSS
		Logo                template.URL
	}{
		Data:                "data goes here",
		Expenses:            []string{"Lysol Aerosol", "Windex Window Cleaner", "Toilet Bowl Cleaner"},
		BootstrapStylesheet: tools.GetStylesheet("./assets/bootstrap.css"),
		Logo:                tools.ConvertImageToBase64("./assets/images/invoicer.jpg"),
	}
	tools.CreatePdf("./templates/test.html", "./storage/pdf.pdf", templateData)
}
