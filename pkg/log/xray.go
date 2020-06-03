package log

import (
	"fmt"

	"github.com/aws/aws-xray-sdk-go/xraylog"
)

// XrayLogger is
type XrayLogger struct{}

// Log is
func (l *XrayLogger) Log(ll xraylog.LogLevel, msg fmt.Stringer) {

	switch ll {
	case xraylog.LogLevelDebug:
		// noisy at debug
		//logger.Debug(msg.String())
	case xraylog.LogLevelInfo:
		logger.Info(msg.String())
	case xraylog.LogLevelWarn:
		logger.Warn(msg.String())
	case xraylog.LogLevelError:
		logger.Error(msg.String())
	}
}
