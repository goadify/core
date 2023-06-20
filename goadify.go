package goadify

import (
	"github.com/goadify/goadify/interfaces"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Goadify struct {
	logger          interfaces.Logger
	config          *Config
	ordinaryModules []interfaces.OrdinaryModule
	extendedModules []interfaces.ExtendedModule
}

func (g *Goadify) fillDefaults() {
	g.logger = logrus.New()
	g.config = new(Config)
}

func (g *Goadify) HttpHandler() (http.Handler, error) {
	mm, err := newModuleMaster(
		g.logger,
		g.config.isDevModeEnabled,
		g.ordinaryModules,
		g.extendedModules,
	)

	if err != nil {
		return nil, err
	}

	return mm.HttpHandler(), nil
}

func New(options ...Option) *Goadify {
	g := new(Goadify)

	g.fillDefaults()

	for _, option := range options {
		option(g)
	}

	return g
}
