package logger

import (
	"cdk-app-template/internal/pkg/constants"
	errors "cdk-app-template/internal/pkg/error"
	"context"
	"go.uber.org/zap"
)

func Create() (*zap.Logger, error) {
	logger, err := zap.NewDevelopment()
	return logger, errors.Wrap("error creating logger", err)
}

func FromContext(ctx context.Context) *zap.Logger {
	return ctx.Value(constants.CTX_LOGGER).(*zap.Logger)
}
