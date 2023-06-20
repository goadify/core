package interfaces

import (
	"net/http"
)

type DependencyProvider interface {
	Logger() Logger
	DevModeEnabled() bool
}

type ExtendedDependencyProvider interface {
	DependencyProvider

	Modules() []Module
}

type Module interface {
	Name() string

	HttpPrefix() string
	HttpHandler() (http.Handler, error)
}

type OrdinaryModule interface {
	Module

	Init(DependencyProvider) error
}

type ExtendedModule interface {
	Module

	Init(ExtendedDependencyProvider) error
}
