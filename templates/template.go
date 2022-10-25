package templates

import (
	"embed"
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

//go:embed views/*.html
var templateFS embed.FS

type Template struct {
	templates *template.Template
}

func NewTemplate() *Template {
	templates := template.Must(template.New("").Funcs(funcMap()).ParseFS(templateFS, "views/*.html"))
	return &Template{
		templates: templates,
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl := template.Must(t.templates.Clone())
	tmpl = template.Must(tmpl.ParseFS(templateFS, fmt.Sprintf("views/%s", name)))
	return tmpl.ExecuteTemplate(w, name, data)
}

func funcMap() template.FuncMap {
	return template.FuncMap{
		"dateTime": dateTime,
		"noEscape": noEscape,
	}
}
