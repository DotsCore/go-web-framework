package go_web_framework

import (
	"gopkg.in/yaml.v2"
	"os"
)

// This is the main configuration of Go-Web
// You can implement this method if wanna implement more configuration.
// Remember: this struct will be populated by parsing the config.yml file present into the Go-Web main directory.
// You've to implement both to works properly.
type Conf struct {
	Database struct {
		Driver   string `yaml:"driver"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Database string `yaml:"database"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"database"`
	Redis struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"redis"`
	Mongo struct {
		Database string `yaml:"database"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"mongodb"`
	Elastic struct {
		Hosts []string `yaml:"hosts"`
	} `yaml:"elasticsearch"`
	Server struct {
		Name     string `yaml:"name"`
		Port     int    `yaml:"port"`
		Ssl      bool   `yaml:"ssl"`
		SslCert  string `yaml:"sslcert"`
		SslKey   string `yaml:"sslkey"`
		RunUser  string `yaml:"run-user"`
		RunGroup string `yaml:"run-group"`
	} `yaml:"server"`
	App struct {
		Key string `yaml:"key"`
	} `yaml:"app"`
	Mail struct {
		From     string `yaml:"from"`
		Host     string `yaml:"host"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Port     int    `yaml:"port"`
	} `yaml:"mail"`
}

// Get configuration struct by parsing the config.yml file.
func Configuration() (*Conf, error) {
	var conf Conf
	confFile := GetDynamicPath("config.yml")
	c, err := os.Open(confFile)

	if err != nil {
		return nil, err
	}

	decoder := yaml.NewDecoder(c)

	if err := decoder.Decode(&conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
