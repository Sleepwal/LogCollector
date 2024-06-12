package initial

import (
	"github.com/sirupsen/logrus"
)

// InitLogger
// @Description: 初始化日志
func InitLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	return logger
}
