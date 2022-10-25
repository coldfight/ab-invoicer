package invoice_generator

import (
	"github.com/coldfight/ab-invoicer/internal/models"
	"github.com/coldfight/ab-invoicer/internal/tools/pdf_generator"
	templateHelpers "github.com/coldfight/ab-invoicer/internal/tools/template_helpers"
	"html/template"
)

type InvoiceTemplateData struct {
	Invoice             models.Invoice
	BootstrapStylesheet template.CSS
	Fonts               map[string]templateHelpers.FontVariation
}

func NewInvoice(invoice models.Invoice) {
	bootstrapStylesheet := templateHelpers.GetStylesheet("assets/styles/bootstrap.css")
	fontMap := map[string]templateHelpers.FontVariation{
		"Normal": {
			Name:    "fira-code",
			Regular: templateHelpers.ConvertFontToBase64("assets/fonts/FiraCode/fira-code-regular.ttf"),
			Bold:    templateHelpers.ConvertFontToBase64("assets/fonts/FiraCode/fira-code-bold.ttf"),
		},
		"Mono": {
			Name:    "fira-code-mono",
			Regular: templateHelpers.ConvertFontToBase64("assets/fonts/FiraCode/fira-code-regular-mono.ttf"),
			Bold:    templateHelpers.ConvertFontToBase64("assets/fonts/FiraCode/fira-code-bold-mono.ttf"),
		},
	}

	templateData := InvoiceTemplateData{
		Invoice:             invoice,
		BootstrapStylesheet: bootstrapStylesheet,
		Fonts:               fontMap,
	}

	pdf_generator.CreatePdf("invoice.tmpl", "./storage/invoice.pdf", templateData)
}
