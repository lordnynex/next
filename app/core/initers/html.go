package initers

import (
	"github.com/sknv/upsale/app/core/cfg"
	xtemplate "github.com/sknv/upsale/app/lib/html/template"
)

const (
	tplRoot = "templates/"
	tplExt  = ".tpl"
)

var (
	html *xtemplate.HTML
)

func init() {
	html = xtemplate.NewHTML(tplRoot, tplExt, cfg.IsProduction())
}

func GetHTML() *xtemplate.HTML {
	return html
}
