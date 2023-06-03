package goadify

import (
	"github.com/goadify/goadify/interfaces"
	"github.com/goadify/goadify/internal/modules/core"
	"net/http"
)

type Goadify struct {
	logger  interfaces.Logger
	modules []interfaces.Module
	config  Config
}

func (g *Goadify) fillDefaults() {
	g.logger = new(loggerStub)
}

func (g *Goadify) loadInternalModules() {
	var modules []interfaces.Module
	copy(modules, g.modules)

	g.loadOptions([]Option{
		WithModule(core.NewModule(modules)),
	})
}

func New(options ...Option) *Goadify {
	g := new(Goadify)

	g.fillDefaults()
	g.loadOptions(options)
	g.loadInternalModules()

	return g
}

func (g *Goadify) Handler() (http.Handler, error) {
	return g.buildHandler()
}
