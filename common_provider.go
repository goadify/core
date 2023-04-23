package goadify

type CommonProvider interface {
	Logger() Logger
}

type commonProvider struct {
	logger Logger
}

func (cp *commonProvider) Logger() Logger {
	return cp.logger
}

func (g *Goadify) commonProvider() CommonProvider {
	return &commonProvider{
		logger: g.logger,
	}
}
