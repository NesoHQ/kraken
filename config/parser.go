package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ServiceConfig struct {
	ServiceName string   `yaml:"service_name"`
	PackageName string   `yaml:"package_name"`
	Database    Database `yaml:"database"`
	Entities    []Entity `yaml:"entities"`
}

type Database struct {
	Type   string `yaml:"type"`
	Schema string `yaml:"schema"`
}

type Entity struct {
	Name      string  `yaml:"name"`
	TableName string  `yaml:"table_name"`
	Fields    []Field `yaml:"fields"`
}

type Field struct {
	Name       string      `yaml:"name"`
	Type       string      `yaml:"type"`
	Primary    bool        `yaml:"primary,omitempty"`
	Nullable   bool        `yaml:"nullable,omitempty"`
	ForeignKey *ForeignKey `yaml:"foreign_key,omitempty"`
	Values     []string    `yaml:"values,omitempty"`
}

type ForeignKey struct {
	ReferenceEntity string `yaml:"reference_entity"`
	ReferenceField  string `yaml:"reference_field"`
}

func LoadConfig() (*ServiceConfig, error) {
	file, err := os.ReadFile("config/sample_config.yml")
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var cfg ServiceConfig
	if err := yaml.Unmarshal(file, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal yaml: %w", err)
	}

	return &cfg, nil
}
