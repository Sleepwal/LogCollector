package initialize

import (
	"LogCollector/config"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

// InitSystemConfig
// @Description: 初始化SystemConfig
func InitSystemConfig() *config.SystemConfig {
	// 默认值
	systemConfig := &config.SystemConfig{
		Name: "LogCollector V1.0 ServerApp（Default Config）",
		KafKaConfig: config.KafKaConfig{
			Address: []string{"127.0.0.1:9092"},
		},
		EtcdConfig: config.EtcdConfig{
			Endpoints: []string{"127.0.0.1:2379"},
		},
	}

	// 加载配置文件conf/config.yaml
	loadConfig(systemConfig)

	logrus.Info("config init success!")
	return systemConfig
}

// LoadConfig
// @Description: 从config.yaml读取配置参数
// @receiver sg
func loadConfig(systemConfig *config.SystemConfig) {
	// 读取 YAML 文件
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		logrus.Warn("[Global]load config.yaml failed, use default config")
		return
	}

	// 解析 YAML 文件
	err = yaml.Unmarshal(file, systemConfig)
	if err != nil {
		logrus.Warn("[Global]unmarshal config.yaml failed, use default config")
		return
	}
}
