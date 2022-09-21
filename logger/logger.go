package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log = NewLogger()

// Write
// @param level
// @param msg
// @param fields
// @date 2022-09-10 17:36:33
func Write(level zapcore.Level, msg string, fields ...zap.Field) {
	if ce := log.Check(level, msg); ce != nil {
		ce.Write(fields...)
	}
}

// Debug
// @param msg
// @param fields
// @date 2022-09-10 17:36:32
func Debug(msg string, fields ...zap.Field) {
	if ce := log.Check(zap.DebugLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

// Info
// @param msg
// @param fields
// @date 2022-09-10 17:36:31
func Info(msg string, fields ...zap.Field) {
	if ce := log.Check(zap.InfoLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

// Warn
// @param msg
// @param fields
// @date 2022-09-10 17:36:30
func Warn(msg string, fields ...zap.Field) {
	if ce := log.Check(zap.WarnLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

// Error
// @param msg
// @param fields
// @date 2022-09-10 17:36:29
func Error(msg string, fields ...zap.Field) {
	if ce := log.Check(zap.ErrorLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

// Fatal
// @param msg
// @param fields
// @date 2022-09-10 17:36:28
func Fatal(msg string, fields ...zap.Field) {
	if ce := log.Check(zap.FatalLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

// Wrap
// @param core
// @date 2022-09-10 17:36:27
func Wrap(core func(zapcore.Core) zapcore.Core) {
	log.WithOptions(zap.WrapCore(core))
}

// WithError
// @param err
// @date 2022-09-10 17:36:26
func WithError(err error) *zap.Logger {
	if err == nil {
		return log
	}
	return log.With(zap.String("error", err.Error()))
}

// WithAction
// @param action
// @date 2022-09-10 17:36:25
func WithAction(action string) *zap.Logger {
	return log.With(zap.String("action", action))
}

// WithData
// @param data
// @date 2022-09-10 17:36:24
func WithData(data any) *zap.Logger {
	return log.With(zap.Any("data", data))
}

// WithApp
// @param app
// @date 2022-09-10 17:36:23
func WithApp(app string) {
	log = log.WithOptions(zap.Fields(zap.String("app", app)))
}

// AppendData
// @param data
// @date 2022-09-10 17:36:22
func AppendData(data any) zap.Field {
	return zap.Any("data", data)
}

func Fields(fields any) zap.Field {
	return zap.Any("data", fields)
}
