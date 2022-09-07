package config

// SystemBaseConfig 基础的系统设置,用于控制worker和job queue数量
type SystemBaseConfig struct {
	ProjectName string   `yaml:"project_name"`
	MaxWorker   int      `yaml:"max_worker"`
	MaxQueue    int      `yaml:"max_queue"`
	Debug       bool     `yaml:"debug"`
	AllowedHost []string `yaml:"allowed_host"`
	SecretKey   string   `yaml:"secret_key"`
	ApiPrefix   string   `yaml:"api_prefix"`
}
