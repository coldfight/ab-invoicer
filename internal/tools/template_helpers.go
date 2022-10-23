package tools

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/coldfight/ab-invoicer/templates"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

const (
	SavedDateLayout = "Jan 02, 2006"
)

func FullFilePath(relativePath string) string {
	abs, err := filepath.Abs(relativePath)
	if err != nil {
		return ""
	}
	return abs
}

func Currency(num float64) string {
	return fmt.Sprintf("$%.2f", num)
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

func (d *Date) MarshalJSON() ([]byte, error) {
	// @todo: I might want to call the json.Marshal instead of manually appending the `"`
	return []byte(`"` + time.Time(*d).Format(SavedDateLayout) + `"`), nil
}

// Format @todo: Figure out why .Date.Format works for the Date in the Labour struct and not for the InvoiceDate in the Receipt struct
// Probably something to do with pointer receiver, etc
func (d *Date) Format(layout string) string {
	return FormatDate(*d, layout)
}

// FormatDate @todo: Figure out why .FormatDate works for the InvoiceDate in the Receipt struct and not for the Date in the Labour struct
// Probably something to do with non-pointer Date param
func FormatDate(d Date, layout string) string {
	return time.Time(d).Format(layout)
}

func getGlobalTemplateFunctions() template.FuncMap {
	return template.FuncMap{
		"GetAbsPath": FullFilePath,
		"AsCurrency": Currency,
		"FormatDate": FormatDate,
	}
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
