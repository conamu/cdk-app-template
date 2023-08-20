package logger

import (
	"cdk-app-template/internal/pkg/constants"
	"context"
	"go.uber.org/zap"
	"log/slog"
	"os"
)

func Create() *slog.Logger {
	stackName := os.Getenv("STACK_NAME")
	stage := os.Getenv("STAGE")

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.LevelDebug,
		ReplaceAttr: nil,
	})
	logger := slog.New(handler)

	logger = logger.With("App", stackName)
	logger = logger.With("Stage", stage)

	return logger
}

func FromContext(ctx context.Context) *zap.Logger {
	return ctx.Value(constants.CTX_LOGGER).(*zap.Logger)
}
