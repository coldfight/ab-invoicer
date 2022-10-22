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
		Fonts               map[string]tools.FontFamily
	}{
		Data:                "data goes here",
		Expenses:            []string{"Lysol Aerosol", "Windex Window Cleaner", "Toilet Bowl Cleaner"},
		BootstrapStylesheet: tools.GetStylesheet("./assets/bootstrap.css"),
		Logo:                tools.ConvertImageToBase64("./assets/images/invoicer.jpg"),
		Fonts: map[string]tools.FontFamily{
			"Normal": {
				Name:    "fira-code",
				Regular: tools.ConvertFontToBase64("./assets/fonts/FiraCode/fira-code-regular.ttf"),
				Bold:    tools.ConvertFontToBase64("./assets/fonts/FiraCode/fira-code-bold.ttf"),
			},
			"Mono": {
				Name:    "fira-code-mono",
				Regular: tools.ConvertFontToBase64("./assets/fonts/FiraCode/fira-code-regular-mono.ttf"),
				Bold:    tools.ConvertFontToBase64("./assets/fonts/FiraCode/fira-code-bold-mono.ttf"),
			},
		},
	}

	tools.CreatePdf("./templates/test.tmpl", "./storage/pdf.pdf", templateData)
}
