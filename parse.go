package goconf

import (
	"fmt"
	"io"
	"io/ioutil"
)

func Parse(r io.Reader, conf interface{}, unmarshal func([]byte, interface{}) error) error {
	buf, err := ioutil.ReadAll(r)

	if err != nil {
		return fmt.Errorf(errorWrapIO, err)
	}

	err = unmarshal(buf, conf)

	if nil != err {
		return fmt.Errorf(errorWrapDecode, err)
	}

	return err

}
