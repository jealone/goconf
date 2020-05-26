package goconf

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func NewYamlConfig(r io.Reader) (nodes []*YamlConfig, err error) {

	data, err := ioutil.ReadAll(r)

	if nil != err {
		err = fmt.Errorf(errorWrapIO, err)
		return
	}

	chunks := bytes.Split(data, []byte("---"))
	nodes = make([]*YamlConfig, 0)

	for _, chunk := range chunks {
		if c := bytes.TrimSpace(chunk); 0 != len(c) {
			conf := newConfig()
			err = Unmarshal(append(c, '\n'), conf)
			if nil != err {
				nodes = nil
				return
			}
			nodes = append(nodes, conf)
		}
	}
	return
}

func NewConfigFile(path string) (nodes []*YamlConfig, err error) {
	var f *os.File
	f, err = os.Open(path)
	if nil != err {
		err = fmt.Errorf(errorWrapIO, err)
		return
	}
	defer f.Close()

	nodes, err = NewYamlConfig(f)

	return
}
