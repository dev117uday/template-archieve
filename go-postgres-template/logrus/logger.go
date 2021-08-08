package logrus

import (
	"github.com/sirupsen/logrus"
)

// Logit : custom log function with different log levels
func LogIt(logCase int, funcName, errMsg string, err error) {
	switch logCase {
	case LogFatal:
		Log.WithFields(logrus.Fields{
			"fn":  funcName,
			"err": err.Error(),
		}).Fatal(errMsg)

	case LogWarn:
		Log.WithFields(logrus.Fields{
			"fn":  funcName,
			"err": err.Error(),
		}).Warn(errMsg)

	case LogInfo:
		Log.WithFields(logrus.Fields{
			"fn":  funcName,
			"err": err.Error(),
		}).Info(errMsg)
	}
}
