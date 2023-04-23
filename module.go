package goadify

import (
	"github.com/pkg/errors"
	"net/http"
)

type Module interface {
	Name() string

	Init(CommonProvider) error

	HttpPrefix() string
	HttpHandler() http.Handler
}

func WithModule(module Module) Option {
	return func(goadify *Goadify) {
		goadify.modules = append(goadify.modules, module)
	}
}

func (g *Goadify) buildModules() (http.Handler, error) {
	mux := http.NewServeMux()
	cp := g.buildCommonProvider()

	for _, module := range g.modules {

		err := checkModule(module)
		if err != nil {
			return nil, errors.Wrapf(err, "can not build module %s", module.Name())
		}

		err = module.Init(cp)
		if err != nil {
			return nil, errors.Wrapf(err, "can not build module %s", module.Name())
		}

		prefix := module.HttpPrefix()

		mux.Handle(
			prefix+"/",
			http.StripPrefix(prefix, module.HttpHandler()),
		)
	}

	return mux, nil
}

var (
	ErrPrefixNotStartingWithSlash = errors.New("module's http prefix should starts with slash")
	ErrPrefixEndsWithSlash        = errors.New("module's http prefix should NOT ends with slash")
)

func checkModule(m Module) error {
	prefix := m.HttpPrefix()

	if prefix[0] != '/' {
		return ErrPrefixNotStartingWithSlash
	}

	if prefix[len(prefix)-1] == '/' {
		return ErrPrefixEndsWithSlash
	}

	return nil
}
