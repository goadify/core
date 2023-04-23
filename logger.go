package goadify

type Logger interface {
	Debug(args ...any)
	Info(args ...any)
	Print(args ...any)
	Warn(args ...any)
	Error(args ...any)
	Fatal(args ...any)
	Panic(args ...any)

	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Printf(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)
	Panicf(format string, args ...any)
}

func WithLogger(logger Logger) Option {
	return func(goadify *Goadify) {
		goadify.logger = logger
	}
}

type loggerStub struct{}

func (l *loggerStub) Debug(...any) {

}

func (l *loggerStub) Info(...any) {

}

func (l *loggerStub) Print(...any) {

}

func (l *loggerStub) Warn(...any) {

}

func (l *loggerStub) Error(...any) {

}

func (l *loggerStub) Fatal(...any) {

}

func (l *loggerStub) Panic(...any) {

}

func (l *loggerStub) Debugf(string, ...any) {

}

func (l *loggerStub) Infof(string, ...any) {

}

func (l *loggerStub) Printf(string, ...any) {

}

func (l *loggerStub) Warnf(string, ...any) {

}

func (l *loggerStub) Errorf(string, ...any) {

}

func (l *loggerStub) Fatalf(string, ...any) {

}

func (l *loggerStub) Panicf(string, ...any) {

}
