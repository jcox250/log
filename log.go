package log

import (
	"io"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type LeveledLogger interface {
	Info(keyvals ...interface{})
	Debug(keyvals ...interface{})
	Error(keyvals ...interface{})
}

type logger struct {
	log.Logger
	debug bool
}

func NewLeveledLogger(w io.Writer, debug bool) LeveledLogger {
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
