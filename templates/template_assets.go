package templates

import (
	"embed"
)

//go:embed assets
//go:embed invoice.tmpl
var TemplateAssets embed.FS
