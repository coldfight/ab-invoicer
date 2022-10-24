package pdf_generator

import (
	"bytes"
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/coldfight/ab-invoicer/internal/tools/template_helpers"
	"github.com/coldfight/ab-invoicer/templates"
	"html/template"
	"log"
	"os"
	"strconv"
	"time"
)

type RequestPdf struct {
	body string
}

// newRequestPdf - constructor function that returns a new RequestPdf object
func newRequestPdf(body string) *RequestPdf {
	return &RequestPdf{
		body: body,
	}
}

// parseTemplate - retrieves the template from the embedded FS, sends the template data and FuncMap
// and sets the request's body with the generated html's content
func (r *RequestPdf) parseTemplate(templateFileName string, data any) error {
	t, err := template.New(templateFileName).
		Funcs(template_helpers.GetGlobalTemplateFunctions()).
		ParseFS(templates.TemplateAssets, templateFileName)

	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}

// generatePdf - takes the request body as html and creates a new pdf file from it
func (r *RequestPdf) generatePdf(pdfPath string) (bool, error) {
	t := time.Now().Unix()
	fileName := fmt.Sprintf("storage/%s.html", strconv.FormatInt(int64(t), 10))
	err := os.WriteFile(fileName, []byte(r.body), 0644)
	if err != nil {
		panic(err)
	}
	f, err := os.Open(fileName)
	if f != nil {
		defer f.Close()
		defer os.Remove(fileName)
	}
	if err != nil {
		log.Fatal(err)
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	page := wkhtmltopdf.NewPageReader(f)
	page.EnableLocalFileAccess.Set(true)

	pdfg.AddPage(page)
	// To me this converted to width: 1062px; height: 1374px; in css
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeLetter)

	// Set the margins and padding through css
	pdfg.MarginLeft.Set(0)
	pdfg.MarginRight.Set(0)
	pdfg.MarginTop.Set(0)
	pdfg.MarginBottom.Set(0)
	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	err = pdfg.WriteFile(pdfPath)
	if err != nil {
		log.Fatal(err)
	}

	return true, nil
}

// CreatePdf - public facing function to create the pdf
func CreatePdf(templatePath, outPath string, templateData any) {
	r := newRequestPdf("")

	if err := r.parseTemplate(templatePath, templateData); err == nil {
		ok, _ := r.generatePdf(outPath)
		fmt.Println(ok, "pdf generated successfully")
	} else {
		fmt.Println(err)
	}
}
