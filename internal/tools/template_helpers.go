package tools

import (
	"encoding/json"
	"fmt"
	"html/template"
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
