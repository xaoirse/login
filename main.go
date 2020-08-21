package main

import (
	"html/template"
	"io"
	"log"

	"github.com/xaoirse/logbook/model"
	_ "github.com/xaoirse/logbook/model"
	"github.com/xaoirse/logbook/router"

	"github.com/labstack/echo/v4"
)

// Template is
type Template struct {
	templates *template.Template
}

// Render is
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func oldMain() {
	db := model.GetDb()
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalln("Error when closing db:", err)
		}
	}()

	r := router.New(db)

	t := &Template{templates: template.Must(template.ParseGlob("template/*.html"))}
	r.Renderer = t

	r.Logger.Fatal(r.Start(":4000"))
}
