package v1

import (
	"github.com/goadify/goadify/interfaces"
	"github.com/goadify/openapi/core/go/gen"
	"net/http"
)

type Module struct {
	loadedModules []interfaces.Module
}

func (m *Module) Name() string {
	return "core"
}

func (m *Module) Init(_ interfaces.CommonProvider) error {
	return nil
}

func (m *Module) HttpPrefix() string {
	return "/core/v1"
}

func (m *Module) HttpHandler() (http.Handler, error) {
	s, err := gen.NewServer(gen.UnimplementedHandler{})
	if err != nil {
		return nil, err
	}

	return s, nil
}

func NewModule(loadedModules []interfaces.Module) *Module {

	return &Module{
		loadedModules: loadedModules,
	}
}
