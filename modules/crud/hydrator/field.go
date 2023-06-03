package hydrator

import (
	"github.com/goadify/goadify/modules/crud/models"
	"github.com/goadify/openapi/crud/go/gen"
)

func Fields(fields []models.Field) []gen.EntityMappingFieldsItem {
	fs := make([]gen.EntityMappingFieldsItem, len(fields))
	for i := 0; i < len(fields); i++ {
		fs[i] = gen.EntityMappingFieldsItem{
			Name:     fields[i].Name,
			Datatype: gen.EntityMappingFieldsItemDatatype(fields[i].Datatype),
		}
	}

	return fs
}
