package tools

import (
	"html/template"
	"log"
	"os"
)

func GetStylesheet(filepath string) template.CSS {
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return template.CSS(string(data))
}
