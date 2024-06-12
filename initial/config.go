package initial

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

/****************************************
@Author : SleepWalker
@Description: 系统配置
@File : global
****************************************/

// SystemConfig
// @Description: 一些参数通过配置文件读取，由用户进行配置
type SystemConfig struct {
	//======= Sever ==========
	Name     string `yaml:"name"`     // 服务器名称
	FilePath string `yaml:"filePath"` // 要读取的日志文件路径

	//======= Kafka ==========
	KafKaConfig `yaml:"kafka"` // kafka配置
}

// KafKaConfig
// @Description: kafka配置
type KafKaConfig struct {
	Address []string `yaml:"address"` // kafka地址
	Topic   string   `yaml:"topic"`   // kafka topic
}

// InitSystemConfig
// @Description: 初始化SystemConfig
func InitSystemConfig() *SystemConfig {
	// 默认值
	config := &SystemConfig{
		Name:     "LogCollector V1.0 ServerApp（Default Config）",
		FilePath: "./demo/test.log",
		KafKaConfig: KafKaConfig{
			Address: []string{"127.0.0.1:9092"},
		},
	}

	// 加载配置文件conf/config.yaml
	config.LoadConfig(config)

	logrus.Info("config init success!")
	return config
}

// LoadConfig
// @Description: 从config.yaml读取配置参数
// @receiver sg
func (sg *SystemConfig) LoadConfig(config *SystemConfig) {
	// 读取 YAML 文件
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		logrus.Error("[Global]load config.yaml failed, use default config")
		return
		//panic(err)
	}

	// 解析 YAML 文件
	err = yaml.Unmarshal(file, config)
	if err != nil {
		panic(err)
	}
}
