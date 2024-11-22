package log

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	sync.RWMutex
	log *zap.Logger
}

func NewLogger(opts ...Option) Logger {
	optionDefault := Option{
		Level:       DebugLevel,
		Encoding:    JsonEncoding,
		Development: false,
	}

	var opt Option
	if len(opts) == 0 {
		opt = optionDefault
	}

	if opt.Level < 0 || opt.Level > 5 {
		opt.Level = optionDefault.Level
	}

	if opt.Encoding != JsonEncoding {
		opt.Encoding = JsonEncoding
	}

	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(convertToZapLevel(opt.Level)),
		Encoding:    string(opt.Encoding),
		Development: opt.Development,
	}

	if len(opt.Fields) > 0 {
		cfg.InitialFields = opt.Fields
	}

	logger, err := cfg.Build()
	if err != nil {
		panic("it is not possible create a log instance")
	}

	return &zapLogger{
		log: logger,
	}
}

func (l *zapLogger) Close() error {
	return l.log.Sync()
}

func (l *zapLogger) AddFields(fields ...Field) {
	l.Lock()
	defer l.Unlock()
	nfields := make([]zap.Field, 0, len(fields))
	for _, f := range fields {
		nfields = append(nfields, zap.Any(f.Key, f.Value))
	}

	l.log = l.log.With(nfields...)
}

func (l *zapLogger) Error(err error) {
	l.AddFields(Field{
		Key:   "error",
		Value: err,
	})
}

func (l *zapLogger) Log(level LogLevel, msg string, fields ...Field) {
	l.Lock()
	defer l.Unlock()

	mfields := make([]zap.Field, len(fields))
	for _, l := range fields {
		mfields = append(mfields, zap.Any(l.Key, l.Value))
	}

	switch convertToZapLevel(level) {
	case zap.DebugLevel:
		l.log.Debug(msg, mfields...)
	case zap.InfoLevel:
		l.log.Info(msg, mfields...)
	case zap.WarnLevel:
		l.log.Warn(msg, mfields...)
	case zap.ErrorLevel:
		l.log.Error(msg, mfields...)
	case zap.FatalLevel:
		l.log.Fatal(msg, mfields...)
	}
}

func convertToZapLevel(l LogLevel) zapcore.Level {
	switch l {
	case InfoLevel:
		return zapcore.InfoLevel
	case WarnLevel:
		return zapcore.WarnLevel
	case ErrorLevel:
		return zapcore.ErrorLevel
	case PanicLevel:
		return zapcore.PanicLevel
	case FatalLevel:
		return zapcore.FatalLevel
	}
	return zapcore.DebugLevel
}
