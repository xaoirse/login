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
	var newStr string
	for i, c := range str {
		if unicode.IsUpper(c) && i != 0 {
			newStr += "_"
		}
		newStr += string(unicode.ToLower(c))
	}
	return newStr
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
		if field.Name == "id" {
			field.Tag += ` gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
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
		},
	)
}

// Defining mutation function
func mutateHook(b *modelgen.ModelBuild) *modelgen.ModelBuild {
	for _, model := range b.Models {
		addGormTags(model)
		addGormFields(model)
	}
	return b
}
