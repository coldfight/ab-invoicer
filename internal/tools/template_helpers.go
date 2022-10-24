package tools

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/coldfight/ab-invoicer/templates"
	"html/template"
	"log"
	"net/http"
	"time"
)

const (
	SavedDateLayout = "Jan 02, 2006"
)

// asCurrency - converts a float to a currency string
func asCurrency(num float64) string {
	return fmt.Sprintf("$%.2f", num)
}

// getGlobalTemplateFunctions - returns a FuncMap of global template functions
func getGlobalTemplateFunctions() template.FuncMap {
	return template.FuncMap{
		"AsCurrency": asCurrency,
	}
}

type Date time.Time

func (d *Date) UnmarshalJSON(bytes []byte) error {
	var v interface{}
	if err := json.Unmarshal(bytes, &v); err != nil {
		return err
	}

	t, err := time.Parse(SavedDateLayout, v.(string))
	if err != nil {
		return err
	}

	*d = Date(t)
	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d).Format(SavedDateLayout))
}

func (d Date) Format(layout string) string {
	return time.Time(d).Format(layout)
}

type Font struct {
	Src template.URL
	Url template.URL
}

type FontFamily struct {
	Name    string
	Bold    Font
	Regular Font
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func GetStylesheet(filepath string) template.CSS {
	data, err := templates.TemplateAssets.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return template.CSS(string(data))
}

func ConvertImageToBase64(filepath string) template.URL {
	// Read the entire file into a byte slice
	b, err := templates.TemplateAssets.ReadFile(filepath)

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

func ConvertFontToBase64(filepath string) Font {
	// Read the entire file into a byte slice
	b, err := templates.TemplateAssets.ReadFile(filepath)

	if err != nil {
		log.Fatal(err)
	}

	var base64EncodingSrc string
	var base64EncodingUrl string

	base64EncodingSrc += "data:application/font-woff2;charset=utf-8;base64,"
	base64EncodingUrl += "data:application/font-woff;charset=utf-8;base64,"

	// Append the base64 encoded output
	base64EncodingSrc += toBase64(b)
	base64EncodingUrl += toBase64(b)

	return Font{
		Src: template.URL(base64EncodingSrc),
		Url: template.URL(base64EncodingUrl),
	}
}
