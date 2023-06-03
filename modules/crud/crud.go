package crud

import (
	"github.com/goadify/goadify/interfaces"
	crudGen "github.com/goadify/openapi/crud/go/gen"
	"net/http"
)

type Module struct {
	logger             interfaces.Logger
	entities           []Entity
	repositories       []Repository
	entityRepositories map[string]Repository

	isDevMod bool
}

func (m *Module) Name() string {
	return "crud"
}

func (m *Module) Init(provider interfaces.CommonProvider) error {
	m.logger = provider.Logger()
	m.isDevMod = provider.IsDevMode()
	return nil
}

func (m *Module) HttpPrefix() string {
	return "/modules/crud/v1"
}

func (m *Module) HttpHandler() (http.Handler, error) {
	em := newEntityMaster(m.entities, m.repositories, m.entityRepositories, m.logger)

	hh := newHttpHandler(em, m.isDevMod)

	srv, err := crudGen.NewServer(hh)
	if err != nil {
		return nil, err
	}

	return srv, nil
}

func NewModule(options ...Option) *Module {
	m := new(Module)

	for _, option := range options {
		option(m)
	}

	return m
}
