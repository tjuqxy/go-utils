package utils

import (
	"io/ioutil"

    "launchpad.net/goyaml"
)

func LoadConf(filename string) (map[string]interface{}, error) {
	c := make(map[string]interface{})
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = goyaml.Unmarshal(content, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
