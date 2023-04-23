package goadify

import "net/http"

type Goadify struct {
	logger  Logger
	modules []Module
}

func (g *Goadify) fillDefaults() {
	g.logger = new(loggerStub)
}

func New(options ...Option) *Goadify {
	g := new(Goadify)

	g.loadOptions(options)

	return g
}

func (g *Goadify) Handler() (http.Handler, error) {
	return g.buildModules()
}
