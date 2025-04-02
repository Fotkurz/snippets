package logging

import (
	"log/slog"
	"os"
	"sync"
)

type Logger struct {
	log *slog.Logger
	h   slog.Handler
}

var Log Logger

func NewLogger(opts ...Option) Logger {

	s := sync.Once{}
	s.Do(func() {
		l := &Logger{}
		for _, o := range opts {
			o(l)
		}
		l.log = slog.New(l.h)
		Log = *l
	})

	return Log
}

// Debug implements Logger.
func (l Logger) Debug(msg string, args ...any) {
	l.log.Debug(msg, args...)
}

// Error implements Logger.
func (l Logger) Error(msg string, args ...any) {
	l.log.Error(msg, args...)
}

// Info implements Logger.
func (l Logger) Info(msg string, args ...any) {
	l.log.Info(msg, args...)
}

// Info implements Logger.
func (l Logger) Warn(msg string, args ...any) {
	l.log.Warn(msg, args...)
}

type Option func(l *Logger)

func WithJsonHandler() Option {
	return func(l *Logger) {
		if l.h != nil {
			l.log.Error("tried to set log handler more than once")
			return
		}
		l.h = slog.NewJSONHandler(os.Stdout, nil)
	}
}
