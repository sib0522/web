package forms

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// HTMLレンダリング
func InitRenderer(e *echo.Echo) {
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("web/views/*.html")),
	}
	e.Renderer = renderer
}
