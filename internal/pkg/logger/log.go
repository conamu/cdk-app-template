package logger

import (
	"cdk-app-template/internal/pkg/constants"
	"context"
	"github.com/spf13/viper"
	"log/slog"
	"os"
)

const (
	DEBUG = "debug"
	INFO  = "info"
	WARN  = "warn"
	ERROR = "error"
)

func Create() *slog.Logger {

	stackName := os.Getenv("STACK_NAME")
	stage := os.Getenv("STAGE")
	logLevel := viper.GetString("log-level")

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource:   true,
		Level:       getLogLevel(logLevel),
		ReplaceAttr: nil,
	})
	logger := slog.New(handler)

	logger = logger.With("App", stackName)
	logger = logger.With("Stage", stage)

	return logger
}

func FromContext(ctx context.Context) *slog.Logger {
	return ctx.Value(constants.CTX_LOGGER).(*slog.Logger)
}

func getLogLevel(level string) slog.Level {
	switch level {
	case DEBUG:
		return slog.LevelDebug
	case INFO:
		return slog.LevelInfo
	case WARN:
		return slog.LevelWarn
	case ERROR:
		return slog.LevelError
	}
	return slog.LevelDebug
}
