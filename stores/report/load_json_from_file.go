package main

import (
	"encoding/json"
	"io/ioutil"
)

func loadJsonFile(filename string, output interface{}) (err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	if err = json.Unmarshal(data, output); err != nil {
		return
	}
	return nil
}
