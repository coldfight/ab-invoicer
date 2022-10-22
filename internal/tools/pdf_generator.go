package tools

import (
	"bytes"
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"html/template"
	"log"
	"os"
	"strconv"
	"time"
)

type RequestPdf struct {
	body string
}

func newRequestPdf(body string) *RequestPdf {
	return &RequestPdf{
		body: body,
	}
}

func (r *RequestPdf) parseTemplate(templateFileName string, data any) error {
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
		//defer os.Remove(fileName) // @todo: Uncomment this
	}
	if err != nil {
		log.Fatal(err)
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(f))
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeLetter)
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

func CreatePdf(templatePath, outPath string, templateData any) {
	r := newRequestPdf("")

	if err := r.parseTemplate(templatePath, templateData); err == nil {
		ok, _ := r.generatePdf(outPath)
		fmt.Println(ok, "pdf generated successfully")
	} else {
		fmt.Println(err)
	}
}
