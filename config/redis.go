package config

// RedisConfig redis基础配置
type RedisConfig struct {
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	Password  string `yaml:"password"`
	DB        int    `yaml:"db"`
	Timeout   int    `yaml:"timeout"`
	MaxActive int    `yaml:"max_active"`
	MaxIdle   int    `yaml:"max_idle"`
	PoolSize  int    `yaml:"pool_size"`
}
