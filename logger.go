// Logger is an easy way to add structured logs to your project
// Based on uber-go/zap
package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Settings struct {
	// DevMode when enabled:
	// - log messages are not wrapped in JSON, easier to read
	// - every log message is annotated with the calling function's file name and line number
	// - if stacktraces are enabled, they're captured for Warn level and above (Error and above otherwise)
	// - log level is Debug (Info otherwise)
	DevMode bool

	// StackTracedErrors when enabled, adds stack traces to each log message of level Error and above
	StackTracedErrors bool
}

var logger *zap.Logger
var level zap.AtomicLevel

func init() {
	// production mode by default
	Init(Settings{})
}

func Init(settings Settings) {
	var config zap.Config
	if settings.DevMode {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()
		config.DisableCaller = true
	}
	level = config.Level

	if !settings.StackTracedErrors {
		config.DisableStacktrace = true
	}

	var err error
	logger, err = config.Build()
	if err != nil {
		log.Fatal("logger init:", err)
	}
}

func SetLevel(l zapcore.Level) {
	level.SetLevel(l)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}
