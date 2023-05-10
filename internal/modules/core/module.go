package core

import (
	"github.com/goadify/goadify/types"
	"github.com/goadify/openapi/core/go/gen"
	"net/http"
)

type Module struct {
	loadedModules []types.Module
}

func (m *Module) Name() string {
	return "core"
}

func (m *Module) Init(_ types.CommonProvider) error {
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

func NewModule(loadedModules []types.Module) *Module {
	var modules []types.Module
	copy(modules, loadedModules)

	return &Module{
		loadedModules: modules,
	}
}
