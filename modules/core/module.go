package core

import (
	"github.com/goadify/goadify/interfaces"
	"github.com/goadify/goadify/modules/core/models"
	"github.com/goadify/openapi/core/go/gen"
	"net/http"
)

type Module struct {
	modules       []interfaces.Module
	modulesModels []models.Module
}

func (m *Module) Name() string {
	return "core"
}

func (m *Module) HttpPrefix() string {
	return "/core/v1"
}

func (m *Module) HttpHandler() (http.Handler, error) {
	hh := newHttpHandler(m.modulesModels)

	return gen.NewServer(hh)
}

func (m *Module) prepare() {
	m.modulesModels = make([]models.Module, len(m.modules))

	for i := 0; i < len(m.modules); i++ {

		m.modulesModels[i] = models.Module{
			Name:     m.modules[i].Name(),
			BasePath: m.modules[i].HttpPrefix(),
		}

	}
}

func (m *Module) Init(provider interfaces.ExtendedDependencyProvider) error {
	m.modules = provider.Modules()
	m.prepare()
	return nil
}

func New() *Module {
	return new(Module)
}
