package goconf

import (
	"gopkg.in/yaml.v3"
)

type YamlConfig struct {
	Version  string    `yaml:"version"`
	Kind     string    `yaml:"kind"`
	Metadata yaml.Node `yaml:"metadata"`
	Spec     yaml.Node `yaml:"spec"`
}

func (c *YamlConfig) GetVersion() string {
	return c.Version
}

func (c *YamlConfig) GetKind() string {
	return c.Kind
}

func (c *YamlConfig) GetMetadata() *yaml.Node {
	return &c.Metadata
}

func (c *YamlConfig) GetSpec() *yaml.Node {
	return &c.Metadata
}

func Unmarshal(buf []byte, c interface{}) error {
	return yaml.Unmarshal(buf, c)
}
