package hook

import (
	"fmt"
	"go/types"
	"os"
	"strings"
	"unicode"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin/modelgen"
)

func init() {
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
}

func camel2Snacke(str string) string {
	for i, c := range str {
		if i == 0 {
			str = strings.Replace(str, string(c), string(unicode.ToLower(c)), 1)
			continue
		}
		if unicode.IsUpper(c) {
			str = strings.Replace(str, string(c), "_"+string(unicode.ToLower(c)), 1)
		}
	}
	return str
}

func addM2mTag(model *modelgen.Object, field *modelgen.Field) {
	str := strings.Split(field.Type.String(), ".")
	typeOfField := str[len(str)-1]
	var m2mName string
	if model.Name > typeOfField {
		m2mName = model.Name + typeOfField
	} else {
		m2mName = typeOfField + model.Name
	}
	field.Tag += ` gorm:"many2many:` + camel2Snacke(m2mName) + `s"`
}

func addGormTags(model *modelgen.Object) {
	for _, field := range model.Fields {
		if strings.HasPrefix(field.Type.String(), "[]") {
			addM2mTag(model, field)
		}
		if field.Name == "ID" {
			field.Tag += ` gorm:"primary_key"`
		}
	}
}

func addGormFields(model *modelgen.Object) {
	// TODO add gorm.Model fields
	var cfg config.Config
	typ := types.NewNamed(
		types.NewTypeName(0, cfg.Model.Pkg(), "time.Time", nil),
		nil,
		nil,
	)
	typP := types.NewNamed(
		types.NewTypeName(0, cfg.Model.Pkg(), "*time.Time", nil),
		nil,
		nil,
	)
	model.Fields = append(model.Fields,
		&modelgen.Field{
			Name:        "CreatedAt",
			Type:        typ,
			Description: "gorm.Model",
		},
		&modelgen.Field{
			Name: "UpdatedAt",
			Type: typ,
		},
		&modelgen.Field{
			Name: "DeletedAt",
			Type: typP,
			Tag:  `sql:"index"`,
		})
	/*
	   ID        uint `gorm:"primary_key"`
	   	CreatedAt time.Time
	   	UpdatedAt time.Time
	   	DeletedAt *time.Time `sql:"index"`
	*/
}

// Defining mutation function
func mutateHook(b *modelgen.ModelBuild) *modelgen.ModelBuild {
	for _, model := range b.Models {
		addGormTags(model)
		addGormFields(model)
	}
	return b
}