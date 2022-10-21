package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type RequestPdf struct {
	body string
}

func NewRequestPdf(body string) *RequestPdf {
	return &RequestPdf{
		body: body,
	}
}

func (r *RequestPdf) ParseTemplate(templateFileName string, data any) error {
	t, err := template.ParseFiles(templateFileName)
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

func (r *RequestPdf) GeneratePdf(pdfPath string) (bool, error) {
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

	pdfg.AddPage(wkhtmltopdf.NewPageReader(f))
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
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

func readBootstrapStylesheet() template.CSS {
	data, err := os.ReadFile("./assets/bootstrap.css")
	if err != nil {
		log.Fatal(err)
	}

	return template.CSS(string(data))
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func convertImageToBase64() template.URL {
	// Read the entire file into a byte slice
	b, err := os.ReadFile("./assets/images/invoicer.jpg")
	if err != nil {
		log.Fatal(err)
	}

	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(b)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += toBase64(b)

	return template.URL(base64Encoding)
}

func main() {
	r := NewRequestPdf("")

	templatePath := "./templates/test.html"
	outPath := "./storage/pdf.pdf"

	templateData := struct {
		Data                string
		Expenses            []string
		BootstrapStylesheet template.CSS
		Logo                template.URL
	}{
		Data:                "data goes here",
		Expenses:            []string{"Lysol Aerosol", "Windex Window Cleaner", "Toilet Bowl Cleaner"},
		BootstrapStylesheet: readBootstrapStylesheet(),
		Logo:                convertImageToBase64(),
	}

	if err := r.ParseTemplate(templatePath, templateData); err == nil {
		ok, _ := r.GeneratePdf(outPath)
		fmt.Println(ok, "pdf generated successfully")
	} else {
		fmt.Println(err)
	}
}
