package model

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// ConfigLoader 是一個接口，用於加載配置
type ConfigLoader interface {
	LoadConfig(filename string) (*Config, error)
}

// YAMLConfigLoader 實現了 ConfigLoader 接口
type YAMLConfigLoader struct{}

func (y *YAMLConfigLoader) LoadConfig(filename string) (*Config, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling YAML: %w", err)
	}

	return &config, nil
}
