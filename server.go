package main

import (
	"html/template"
	"io"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	_ "github.com/xaoirse/logbook/graph/hook"
	"github.com/xaoirse/logbook/graph/model"
	"github.com/xaoirse/logbook/router"
)

// Template is
type Template struct {
	templates *template.Template
}

// Render is
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

const defaultPort = "4000"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := model.GetDb()
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalln("Error when closing db:", err)
		}
	}()

	r := router.New(db)

	t := &Template{templates: template.Must(template.ParseGlob("template/*.html"))}
	r.Renderer = t

	r.Logger.Fatal(r.Start(":" + defaultPort))
}
