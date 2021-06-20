package logger

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go-react-mongo/internal"
	"go-react-mongo/internal/config"
	"regexp"
)

var logger *logrus.Logger
type level uint32
const correlationId = "correlation-id"
var CorrelationId = "00000.00000"

func init() {
	logger = logrus.New()
	formatter := logrus.JSONFormatter{}
	logger.SetFormatter(&formatter)
	logger.SetLevel(GetLogLevel(viper.GetString(config.LogLevel)))
}

func Fatalf(ctx context.Context, format string, args ...interface{}) {
	logger.WithField(CorrelationId, internal.GetContextValue(ctx, internal.ContextKeyCorrelationID)).Fatalf(format, args...)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	logger.WithField(CorrelationId, internal.GetContextValue(ctx, internal.ContextKeyCorrelationID)).Infof(format, args...)
}

func Info(ctx context.Context, msg string) {
	logger.WithField(CorrelationId, internal.GetContextValue(ctx, internal.ContextKeyCorrelationID)).Info(msg)
}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	formattedError := escapeString(format, args...)
	logger.WithField(CorrelationId, internal.GetContextValue(ctx, internal.ContextKeyCorrelationID)).Debug(formattedError)
}

func Warnf(ctx context.Context, format string, args ...interface{}) {
	logger.WithField(CorrelationId, internal.GetContextValue(ctx, internal.ContextKeyCorrelationID)).Warnf(format, args...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	formattedError := escapeString(format, args...)
	logger.WithField(CorrelationId, internal.GetContextValue(ctx, internal.ContextKeyCorrelationID)).Error(formattedError)
}

func escapeString(format string, args ...interface{}) string {
	errorMessage := fmt.Sprintf(format, args...)
	re := regexp.MustCompile(`(\n)|(\r\n)`)
	formattedError := re.ReplaceAllString(errorMessage, "\\n ")
	return formattedError
}


func GetLogLevel(level string) logrus.Level {
	switch level {
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warning":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	default:
		return logrus.InfoLevel
	}
}