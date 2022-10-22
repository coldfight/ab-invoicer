package tools

import (
	"encoding/base64"
	"html/template"
	"log"
	"net/http"
	"os"
)

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
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
