package goadify

type Option func(*Goadify)

func (g *Goadify) loadOptions(options []Option) {
	for _, option := range options {
		option(g)
	}
}
