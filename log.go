package log

var (
	defaultLogger Logger = NewLogger()
)

type Logger interface {
	Close() error
	AddFields(fields ...Field)
	Error(error)
	Log(level LogLevel, msg string, fields ...Field)
}
