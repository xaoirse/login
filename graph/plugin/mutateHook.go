package plugin

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

func mutateHook(b *modelgen.ModelBuild) *modelgen.ModelBuild {
	for _, model := range b.Models {
		addGormTags(model)
		addGormFields(model)
	}
	return b
}

func addGormTags(model *modelgen.Object) {
	for _, field := range model.Fields {
		// many2many tag
		if strings.HasPrefix(field.Type.String(), "[]") {
			addM2mTag(model, field)
		}
		// ID tag
		if field.Name == "id" {
			field.Tag += ` gorm:"primary_key"`
			// Why?
			// ` gorm:"primary_key;type:uuid;default:uuid_generate_v4()`
		}
	}
}

type name string

func addM2mTag(model *modelgen.Object, field *modelgen.Field) {
	// str := strings.Split(field.Type.String(), ".")
	// typeOfField := str[len(str)-1]
	typeOfField := field.Type.String()[strings.LastIndex(field.Type.String(), ".")+1:]
	var m2mName name
	if model.Name > typeOfField {
		m2mName = name(model.Name + typeOfField)
	} else {
		m2mName = name(typeOfField + model.Name)
	}
	field.Tag += fmt.Sprintf(` gorm:"many2many:%ss`, m2mName.snackeCase())
}

func (str *name) snackeCase() name {
	var newStr name
	for i, c := range *str {
		if unicode.IsUpper(c) && i != 0 {
			newStr += "_"
		}
		newStr += name(unicode.ToLower(c))
	}
	return newStr
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
		},
	)
}
