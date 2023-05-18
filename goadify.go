package goadify

import (
	"github.com/goadify/goadify/interfaces"
	"github.com/goadify/goadify/internal/modules/core"
	"net/http"
)

type Goadify struct {
	logger  interfaces.Logger
	modules []interfaces.Module
}

func (g *Goadify) fillDefaults() {
	g.logger = new(loggerStub)
}

func (g *Goadify) loadCore() {
	g.loadOptions([]Option{
		WithModule(core.NewModule(g.modules)),
	})
}

func New(options ...Option) *Goadify {
	g := new(Goadify)

	g.fillDefaults()
	g.loadOptions(options)
	g.loadCore()

	return g
}

func (g *Goadify) Handler() (http.Handler, error) {
	return g.buildModules()
}
