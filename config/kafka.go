package config

// KafKaConfig
// @Description: kafka配置
type KafKaConfig struct {
	Address  []string `yaml:"address"`   // kafka地址
	Topic    string   `yaml:"topic"`     // kafka topic
	ChanSize int      `yaml:"chan_size"` // kafka channel大小
}
