package goconf

import (
	"gopkg.in/yaml.v3"
)

type YamlConfig struct {
	Version  string    `yaml:"version"`
	Kind     string    `yaml:"kind"`
	Key      string    `yaml:"key"`
	Metadata yaml.Node `yaml:"metadata"`
	Spec     yaml.Node `yaml:"spec"`
	RawByte  []byte
}

func (c *YamlConfig) GetVersion() string {
	return c.Version
}

func (c *YamlConfig) GetKind() string {
	return c.Kind
}

func (c *YamlConfig) GetKey() string {
	return c.Key
}

func (c *YamlConfig) GetMetadata() *yaml.Node {
	return &c.Metadata
}

func (c *YamlConfig) GetSpec() *yaml.Node {
	return &c.Spec
}

func Unmarshal(buf []byte, c interface{}) error {
	return yaml.Unmarshal(buf, c)
}
