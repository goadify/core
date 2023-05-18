package crud

import (
	"github.com/goadify/goadify/interfaces"
	"net/http"
)

type Module struct {
	logger             interfaces.Logger
	entities           []Entity
	repositories       []Repository
	entityRepositories map[string]Repository
}

func (m *Module) Name() string {
	return "crud"
}

func (m *Module) Init(provider interfaces.CommonProvider) error {
	m.logger = provider.Logger()
	return nil
}

func (m *Module) HttpPrefix() string {
	return "/modules/crud/v1"
}

func (m *Module) HttpHandler() (http.Handler, error) {
	//TODO implement me
	panic("implement me")
}

func NewModule(options ...Option) *Module {
	m := new(Module)

	for _, option := range options {
		option(m)
	}

	return m
}
