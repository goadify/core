package goadify

import "github.com/goadify/goadify/interfaces"

func WithLogger(logger interfaces.Logger) Option {
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
