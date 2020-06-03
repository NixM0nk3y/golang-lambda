package log

import (
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"go.uber.org/zap"
)

// Default logger of the system.
var sdkLogger *AWSLogger

var awsLogLevelSeverity = map[string]aws.LogLevelType{
	"DEBUG":     aws.LogDebugWithHTTPBody,
	"INFO":      aws.LogOff,
	"WARNING":   aws.LogOff,
	"ERROR":     aws.LogOff,
	"CRITICAL":  aws.LogOff,
	"ALERT":     aws.LogOff,
	"EMERGENCY": aws.LogOff,
}

// AWSLogger is
type AWSLogger struct {
}

// AWSLevel is
func AWSLevel() *aws.LogLevelType {

	logLevel := strings.ToUpper(os.Getenv("LOG_LEVEL"))

	if logLevel == "" {
		logLevel = "INFO"
	}

	return aws.LogLevel(awsLogLevelSeverity[logLevel])
}

// Log is
func (l *AWSLogger) Log(args ...interface{}) {
	logger.Debug("awslog", zap.Reflect("output", args))

}
