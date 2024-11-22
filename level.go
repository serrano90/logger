package log

type LogLevel int

var strLevel = []string{"debug", "info", "warn", "error", "panic", "fatal"}

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

func (l LogLevel) String() string {
	if l < 0 || l > 5 {
		return strLevel[0]
	}
	return strLevel[l]
}

func StringToLogLevel(level string) LogLevel {
	for i, v := range strLevel {
		if v == level {
			return LogLevel(i)
		}
	}
	return DebugLevel
}
