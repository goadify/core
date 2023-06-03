package hydrator

import (
	"github.com/goadify/goadify/modules/crud/models"
	"github.com/goadify/openapi/crud/go/gen"
)

func EntityMapping(mapping models.EntityMapping) gen.EntityMapping {
	return gen.EntityMapping{Name: mapping.Name, Fields: Fields(mapping.Fields)}
}

func EntityMappings(mappings []models.EntityMapping) []gen.EntityMapping {
	ems := make([]gen.EntityMapping, len(mappings))

	for i := 0; i < len(mappings); i++ {
		ems[i] = EntityMapping(mappings[i])
	}

	return ems
}
