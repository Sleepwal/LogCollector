package global

import (
	"LogCollector/config"
	"LogCollector/initialize"
	"github.com/sirupsen/logrus"
)

var (
	LOG    *logrus.Logger       // 日志
	CONFIG *config.SystemConfig // 系统配置
)

func init() {
	LOG = initialize.Logger()
	CONFIG = initialize.InitSystemConfig()
}
