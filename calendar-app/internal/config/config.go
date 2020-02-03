package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config struct for configuring app
type Config struct {
	HTTPListen string `yaml:"http_listen"`
	LogFile    string `yaml:"log_file"`
	LogLevel   string `yaml:"log_level"`
}

// FromFile fill Config struct with config from given path
func (c *Config) FromFile(path string) error {
	yamlFile, err := ioutil.ReadFile(path)

	if err != nil {
		return fmt.Errorf("reading config file error: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, c)

	if err != nil {
		return fmt.Errorf("unmarshaling error: %v", err)
	}

	return nil
}
