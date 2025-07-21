package logger

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ctxKey string

const (
	traceIDKey ctxKey = "traceID"
)

var (
	log *zap.Logger
)

func init() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	
	var err error
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func NewContextWithTraceID(ctx context.Context) context.Context {
	return context.WithValue(ctx, traceIDKey, uuid.New().String())
}

func TraceIDFromContext(ctx context.Context) string {
	if traceID, ok := ctx.Value(traceIDKey).(string); ok {
		return traceID
	}
	return ""
}

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, zap.String("traceID", TraceIDFromContext(ctx)))
	log.Info(msg, fields...)
}

func Error(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, zap.String("traceID", TraceIDFromContext(ctx)))
	log.Error(msg, fields...)
}

func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, zap.String("traceID", TraceIDFromContext(ctx)))
	log.Warn(msg, fields...)
}

func Debug(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, zap.String("traceID", TraceIDFromContext(ctx)))
	log.Debug(msg, fields...)
}