package main

import (
	"html/template"
	"io"

	"github.com/xaoirse/logbook/router"

	"github.com/labstack/echo"
)

// Template is
type Template struct {
	templates *template.Template
}

// Render is
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	r := router.New()

	t := &Template{templates: template.Must(template.ParseGlob("views/*.html"))}
	r.Renderer = t

	r.Logger.Fatal(r.Start(":4000"))
}
