package forms

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func getRenderer(fileName string) *TemplateRenderer {
	renderer := &TemplateRenderer{
		//templates: template.Must(goingtpl.ParseFile(fileName)),
		templates: template.Must(template.ParseGlob("./views/*.html")),
	}
	return renderer
}

func SetRenderer(e *echo.Echo, fileName string) {
	e.Renderer = getRenderer(fileName)
}
