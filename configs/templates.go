package configs

import "html/template"

// TPL ...
var TPL *template.Template

func init() {
	funcMap := template.FuncMap{
		"incIdx": func(i int) int {
			return i + 1
		},
	}

	TPL = template.Must(template.New("").Funcs(funcMap).ParseGlob("../web/templates/*.html"))
}
