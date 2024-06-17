package main

import "io/ioutil"

func writeFile(fileName, data string) {
	err := ioutil.WriteFile(fileName, []byte(data), 0600)
	if err != nil {
		panic(err)
	}
}
