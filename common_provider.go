package goadify

import "github.com/goadify/goadify/types"

type commonProvider struct {
	logger types.Logger
}

func (cp *commonProvider) Logger() types.Logger {
	return cp.logger
}

func (g *Goadify) commonProvider() types.CommonProvider {
	return &commonProvider{
		logger: g.logger,
	}
}
