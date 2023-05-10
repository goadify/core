package goadify

import (
	"github.com/goadify/goadify/internal/modules/core"
	"github.com/goadify/goadify/types"
	"net/http"
)

type Goadify struct {
	logger  types.Logger
	modules []types.Module
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
