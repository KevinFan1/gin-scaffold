package config

//LoggerConfig 日志配置
type LoggerConfig struct {
	Path      string `yaml:"path"`
	Prefix    string `yaml:"prefix"`
	Suffix    string `yaml:"suffix"`
	Format    string `yaml:"format"`
	NormalLog string `yaml:"normal_log"`
	ErrorLog  string `yaml:"error_log"`
	MaxSize   int    `yaml:"max_size"`
	Backups   int    `yaml:"backups"`
	Age       int    `yaml:"age"`
	Compress  bool   `yaml:"compress"`
}
