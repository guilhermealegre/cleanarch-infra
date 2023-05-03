package logger

import infra "github.com/guilhermealegre/cleanarch-infra"

type logger struct {
}

func New() infra.ILogger {
	return &logger{}
}

func (l *logger) Init() error {
	return nil
}
