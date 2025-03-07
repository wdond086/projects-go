package task_cli

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
	"os"
)

type ctxKey string

const (
	transactionIDKey ctxKey = "transactionID"
	loggerKey        ctxKey = "logger"
)

func generateTransactionID() string {
	return fmt.Sprintf("trans-%d", rand.Int63())
}

// Injects the logger into the context
func WithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

// Injects the transaction ID into the context
func WithTransactionId(ctx context.Context) context.Context {
	return context.WithValue(ctx, transactionIDKey, generateTransactionID())
}

func NewLogger() *slog.Logger {
	// Console Handlers
	stdoutConsoleHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	})
	return slog.New(stdoutConsoleHandler)
}

// Add the transactionId in the context to the logger
func FromContext(ctx context.Context) *slog.Logger {
	logger := slog.Default()
	if loggerFromContext, ok := ctx.Value(loggerKey).(*slog.Logger); ok {
		logger = loggerFromContext
	}
	if transID, ok := ctx.Value(transactionIDKey).(string); ok {
		logger = logger.With(slog.String(string(transactionIDKey), transID))
	}
	return logger
}
