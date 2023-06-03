package interfaces

type CommonProvider interface {
	Logger() Logger
	IsDevMode() bool
}
