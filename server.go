package main

import (
	"fmt"
	"go/types"
	"html/template"
	"io"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin/modelgen"
	"github.com/labstack/echo/v4"
	"github.com/xaoirse/logbook/model"
	"github.com/xaoirse/logbook/router"
)

// Defining mutation function
func mutateHook(b *modelgen.ModelBuild) *modelgen.ModelBuild {
	for _, model := range b.Models {
		for _, field := range model.Fields {
			if strings.HasPrefix(field.Type.String(), "[]") {
				splitedType := strings.Split(field.Type.String(), ".")
				typeOfField := splitedType[len(splitedType)-1]
				name := typeOfField + model.Name
				if model.Name > typeOfField {
					name = model.Name + typeOfField
				}
				for i, c := range name {
					if i == 0 {
						name = strings.Replace(name, string(c), string(unicode.ToLower(c)), 1)
						continue
					}
					if unicode.IsUpper(c) {
						name = strings.Replace(name, string(c), "_"+string(unicode.ToLower(c)), 1)
					}
				}
				// TODO ID most have gorm tags
				field.Tag += ` gorm:"many2many:` + name + `"`
			}
		}
		// TODO add gorm.Model fields
		typ := types.Typ[types.String].Underlying()
		model.Fields = append(model.Fields, &modelgen.Field{
			Name: "updateAt",
			Type: typ,
		})

	}
	return b
}

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

	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}

	// Attaching the mutation function onto modelgen plugin
	p := modelgen.Plugin{
		MutateHook: mutateHook,
	}

	err = api.Generate(cfg,
		api.NoPlugins(),
		api.AddPlugin(&p),
	)

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
