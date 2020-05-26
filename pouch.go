package goconf

type Decoder interface {
	Decode(interface{}) error
}

func newConfig() *YamlConfig {
	return &YamlConfig{}
}
