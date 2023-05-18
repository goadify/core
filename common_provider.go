package goadify

import "github.com/goadify/goadify/interfaces"

type commonProvider struct {
	logger interfaces.Logger
}

func (cp *commonProvider) Logger() interfaces.Logger {
	return cp.logger
}

func (g *Goadify) commonProvider() interfaces.CommonProvider {
	return &commonProvider{
		logger: g.logger,
	}
}
