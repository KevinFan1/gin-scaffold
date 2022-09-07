package config

import "fmt"

type MySqlConfig struct {
	Username          string `yaml:"username"`
	Password          string `yaml:"password"`
	Host              string `yaml:"host"`
	Port              string `yaml:"port"`
	DBName            string `yaml:"db_ame"`
	ShowLog           bool   `yaml:"show_log"`
	DisableConstraint bool   `yaml:"disable_constraint"`
}

func (config MySqlConfig) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
	)
}
