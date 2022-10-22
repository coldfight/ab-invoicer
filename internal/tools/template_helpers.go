package tools

import (
	"encoding/base64"
	"html/template"
	"log"
	"net/http"
	"os"
)

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
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return template.CSS(string(data))
}

func ConvertImageToBase64(filepath string) template.URL {
	// Read the entire file into a byte slice
	b, err := os.ReadFile(filepath)
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
	b, err := os.ReadFile(filepath)
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
