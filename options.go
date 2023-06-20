package goadify

import "github.com/goadify/goadify/interfaces"

type Option func(g *Goadify)

func WithLogger(logger interfaces.Logger) Option {
	return func(g *Goadify) {
		g.logger = logger
	}
}

func WithModules(modules ...interfaces.Module) Option {
	return func(g *Goadify) {
		for _, module := range modules {
			if om, ok := module.(interfaces.OrdinaryModule); ok {
				g.ordinaryModules = append(g.ordinaryModules, om)
			} else if em, ok := module.(interfaces.ExtendedModule); ok {
				g.extendedModules = append(g.extendedModules, em)
			}
		}
	}
}

func WithConfig(config *Config) Option {
	return func(g *Goadify) {
		g.config = config
	}
}
