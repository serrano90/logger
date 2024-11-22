package log

import "context"

func Init(opts ...Option) {
	defaultLogger = NewLogger(opts...)
}

func Info(msg string, fields ...Field) {
	defaultLogger.Log(InfoLevel, msg, fields...)
}

func InfoCtx(ctx context.Context, msg string, fields ...Field) {
	// Get context fields and join there with the custom fields
	defaultLogger.Log(InfoLevel, msg, fields...)
}

func Warn(msg string, fields ...Field) {
	defaultLogger.Log(WarnLevel, msg, fields...)
}

func WarnCtx(ctx context.Context, msg string, fields ...Field) {
	defaultLogger.Log(InfoLevel, msg, fields...)
}

func Error(msg string, fields ...Field) {
	defaultLogger.Log(ErrorLevel, msg, fields...)
}

func ErrorCtx(ctx context.Context, msg string, fields ...Field) {
	defaultLogger.Log(InfoLevel, msg, fields...)
}

func Debug(msg string, fields ...Field) {
	defaultLogger.Log(DebugLevel, msg, fields...)
}

func DebugCtx(ctx context.Context, msg string, fields ...Field) {
	defaultLogger.Log(InfoLevel, msg, fields...)
}

func Fatal(msg string, fields ...Field) {
	defaultLogger.Log(FatalLevel, msg, fields...)
}

func FatalCtx(ctx context.Context, msg string, fields ...Field) {
	defaultLogger.Log(InfoLevel, msg, fields...)
}

func Panic(msg string, fields ...Field) {
	defaultLogger.Log(PanicLevel, msg, fields...)
}

func PanicCtx(ctx context.Context, msg string, fields ...Field) {
	defaultLogger.Log(InfoLevel, msg, fields...)
}

func AddFields(fields ...Field) {
	defaultLogger.AddFields(fields...)
}

func Close() {
	defaultLogger.Close()
}
