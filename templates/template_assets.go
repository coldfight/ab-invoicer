package templates

import (
	"embed"
)

//go:embed assets
//go:embed receipt.tmpl
var TemplateAssets embed.FS
