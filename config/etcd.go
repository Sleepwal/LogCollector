package config

type EtcdConfig struct {
	Endpoints  []string `yaml:"endpoints"`
	CollectKey string   `yaml:"collect_key"`
}
