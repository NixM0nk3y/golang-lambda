package log

import (
	"context"
	"os"
	"strings"

	"github.com/NixM0nk3y/golang-lambda/pkg/version"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type correlationIDType int

const (
	requestIDKey correlationIDType = iota
	sessionIDKey
)

// Default logger of the system.
var logger *zap.Logger

var logLevelSeverity = map[string]zapcore.Level{
	"DEBUG":     zapcore.DebugLevel,
	"INFO":      zapcore.InfoLevel,
	"WARNING":   zapcore.WarnLevel,
	"ERROR":     zapcore.ErrorLevel,
	"CRITICAL":  zapcore.DPanicLevel,
	"ALERT":     zapcore.PanicLevel,
	"EMERGENCY": zapcore.FatalLevel,
}

func init() {

	buildVersion := version.Version
	buildHash := version.BuildHash
	buildDate := version.BuildDate

	logLevel := strings.ToUpper(os.Getenv("LOG_LEVEL"))

	if logLevel == "" {
		logLevel = "INFO"
	}

	config := zap.NewProductionEncoderConfig()
	//config.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	//	nanos := t.UnixNano()
	//	millis := nanos / int64(time.Millisecond)
	//	enc.AppendInt64(millis)
	//}
	encoder := zapcore.NewJSONEncoder(config)
	atom := zap.NewAtomicLevel()
	defaultLogger := zap.New(zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), atom))

	defer defaultLogger.Sync()

	atom.SetLevel(logLevelSeverity[logLevel])

	logger = defaultLogger.With(zap.String("v", buildVersion), zap.String("bh", buildHash), zap.String("bd", buildDate))
}

// WithRqID returns a context which knows its request ID
func WithRqID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey, requestID)
}

// Logger returns a zap logger with as much context as possible
func Logger(ctx context.Context) *zap.Logger {

	newLogger := logger

	if ctx == nil {
		return newLogger
	}

	if ctxRqID, ok := ctx.Value(requestIDKey).(string); ok {
		newLogger = newLogger.With(zap.String("rqID", ctxRqID))
	}

	return newLogger
}
