package hydrator

import (
	"github.com/goadify/goadify/modules/core/models"
	"github.com/goadify/openapi/core/go/gen"
)

func Module(module models.Module) gen.Module {
	return gen.Module{
		Name:     module.Name,
		BasePath: module.BasePath,
	}
}

func Modules(modules []models.Module) []gen.Module {
	res := make([]gen.Module, len(modules))
	for i := 0; i < len(modules); i++ {
		res[i] = Module(modules[i])
	}

	return res
}
