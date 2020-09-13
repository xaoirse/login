package router

import (
	"io"
	"text/template"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

// Template for echo
type Template struct {
	templates *template.Template
}

// Render for echo
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// Start Echo
func Start(db *gorm.DB, port, secret *string) {

	if *secret == "" {
		*secret = uuid.NewV4().String()
	}

	r := New(db, secret)

	t := &Template{templates: template.Must(template.ParseGlob("template/*.html"))}
	r.Renderer = t

	r.Logger.Fatal(r.Start(":" + *port))
}
