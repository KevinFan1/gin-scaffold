package config

// CeleryConfig celery基本配置，暂时不用
type CeleryConfig struct {
	Broker  string `yaml:"broker"`
	Backend string `yaml:"backend"`
}
