package goadify

import "github.com/goadify/goadify/interfaces"

type commonProvider struct {
	logger    interfaces.Logger
	isDevMode bool
}

func (cp *commonProvider) Logger() interfaces.Logger {
	return cp.logger
}

func (cp *commonProvider) IsDevMode() bool {
	return cp.isDevMode
}

func (g *Goadify) commonProvider() interfaces.CommonProvider {
	return &commonProvider{
		logger:    g.logger,
		isDevMode: g.config.IsDevMode,
	}
}
