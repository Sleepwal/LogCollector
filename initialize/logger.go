package initialize

import (
	"github.com/sirupsen/logrus"
)

// Logger
// @Description: 初始化日志
func Logger() *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	return logger
}
