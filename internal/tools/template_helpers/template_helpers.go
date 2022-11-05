package template_helpers

import (
	"encoding/base64"
	"fmt"
	"github.com/coldfight/ab-invoicer/templates"
	"html/template"
	"log"
	"net/http"
)

// AsCurrency - template function that returns a currency string of a float
func AsCurrency(num float64) string {
	return fmt.Sprintf("$%.2f", num)
}

// PaddedNumber - template function that pads the number with leading 0's
func PaddedNumber(num int) string {
	return fmt.Sprintf("%03d", num)
}

// GetGlobalTemplateFunctions - returns a FuncMap of global template functions
func GetGlobalTemplateFunctions() template.FuncMap {
	return template.FuncMap{
		"AsCurrency":   AsCurrency,
		"PaddedNumber": PaddedNumber,
	}
}

// Font - in css represents the src, and url of the @font-family
type Font struct {
	Src template.URL
	Url template.URL
}

// FontVariation - different font variation for a particular font
type FontVariation struct {
	Name    string
	Bold    Font
	Regular Font
}

// toBase64 - Converts byte slice to base64 string
func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// GetStylesheet - Retrieves the content of a stylesheet and returns as a template.CSS string
// It retrieves the stylesheet from the embedded FS
func GetStylesheet(filepath string) template.CSS {
	data, err := templates.TemplateAssets.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return template.CSS(string(data))
}

// ConvertImageToBase64 - Get the base64 representation of an image and returns it as a template.URL string
// It retrieves the image from the embedded FS
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

// ConvertFontToBase64 - Get the base64 representation of a font file and returns it as a pair of template.URL strings
// It retrieves the font file from the embedded FS
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
