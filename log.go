package loglvl

import (
	"io"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

// Logger is a logger with levels
type Logger interface {
	Info(keyvals ...interface{})
	Debug(keyvals ...interface{})
	Error(keyvals ...interface{})
}

type logger struct {
	log.Logger
	debug bool
}

// NewLogger returns a new leveled logger
func NewLogger(w io.Writer, debug bool) Logger {
	w = log.NewSyncWriter(w)
	kitlogger := log.NewLogfmtLogger(w)
	level.NewFilter(kitlogger, level.AllowAll())
	return &logger{kitlogger, debug}
}

func (l *logger) Info(keyvals ...interface{}) {
	level.Info(l).Log(keyvals...)
}
func (l *logger) Debug(keyvals ...interface{}) {
	if !l.debug {
		return
	}
	level.Debug(l).Log(keyvals...)
}
func (l *logger) Error(keyvals ...interface{}) {
	level.Error(l).Log(keyvals...)
}
