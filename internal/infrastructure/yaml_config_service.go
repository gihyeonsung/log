package infrastructure

import (
	"os"

	"gopkg.in/yaml.v3"
)

type YamlConfigService struct {
	path string
}

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
	Sqlite struct {
		Path string `yaml:"path"`
	} `yaml:"sqlite"`
	Elasticsearch struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"elasticsearch"`
	AuthnService struct {
		Key string `yaml:"key"`
	} `yaml:"authn-service"`
}

func NewYamlConfigService(path string) *YamlConfigService {
	return &YamlConfigService{path: path}
}

func (s *YamlConfigService) Load() (*Config, error) {
	bs, err := os.ReadFile(s.path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(bs, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
