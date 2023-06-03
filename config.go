package goadify

type Config struct {
	IsDevMode bool
}

func WithConfig(config Config) Option {
	return func(goadify *Goadify) {
		goadify.config = config
	}
}
