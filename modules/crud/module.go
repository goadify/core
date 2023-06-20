package crud

import (
	"github.com/goadify/goadify/interfaces"
	"github.com/goadify/goadify/modules/navigation"
	"github.com/goadify/openapi/crud/go/gen"
	"net/http"
)

type Module struct {
	entities       []Entity
	em             *entityMaster
	devModeEnabled bool
}

func (m *Module) Name() string {
	return "crud"
}

func (m *Module) HttpPrefix() string {
	return "/crud/v1"
}

func (m *Module) HttpHandler() (http.Handler, error) {
	hh := newHttpHandler(m.em, m.devModeEnabled)

	return gen.NewServer(hh)
}

func (m *Module) Init(provider interfaces.DependencyProvider) error {
	m.devModeEnabled = provider.DevModeEnabled()
	m.em = newEntityMaster(provider.Logger(), m.entities)
	return nil
}

func (m *Module) Links() []*navigation.Link {
	return m.em.Links()
}

func (m *Module) GroupLinks() map[*navigation.Group][]*navigation.Link {
	return m.em.GroupLinks()
}

func New(entities ...Entity) *Module {
	m := &Module{
		entities: entities,
	}

	return m
}
