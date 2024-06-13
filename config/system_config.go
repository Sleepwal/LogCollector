package config

// SystemConfig
// @Description: 一些参数通过配置文件读取，由用户进行配置
type SystemConfig struct {
	//======= Sever ==========
	Name     string `yaml:"name"`      // 服务器名称
	FilePath string `yaml:"file_path"` // 要读取的日志文件路径

	//======= Kafka ==========
	KafKaConfig `yaml:"kafka"` // kafka配置
	EtcdConfig  `yaml:"etcd"`  // etcd配置
}
