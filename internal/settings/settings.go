package settings

import (
	"code/gin-scaffold/config"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"path"

	"io/ioutil"
)

// Set settings 集合
type Set struct {
	SystemBaseConfig config.SystemBaseConfig `yaml:"sys"`
	RedisConfig      config.RedisConfig      `yaml:"redis"`
	LoggerConfig     config.LoggerConfig     `yaml:"log"`
	CeleryConfig     config.CeleryConfig     `yaml:"celery"`
	MysqlConfig      config.MySqlConfig      `yaml:"mysql"`
}

var Setting = Set{}

func SetUp() {
	projectDir, _ := os.Getwd()
	fileName := path.Join(projectDir, "config", "app.yml")

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("无法读取app yaml文件 file:", err)
	}
	err = yaml.Unmarshal(file, &Setting)
	if err != nil {
		log.Fatal("fail to unmarshal app yaml file:", err)
	}
	if len(Setting.SystemBaseConfig.AllowedHost) == 0 {
		log.Fatal("Allowed Host must not be empty in deployment.")
	}

	log.Printf("读取配置文件%v完成...\n", fileName)
}
