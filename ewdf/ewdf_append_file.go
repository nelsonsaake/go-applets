package main

import "os"

func appendFile(fileName, newData string) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	_, err = f.WriteString(newData)
	if err != nil {
		panic(err)
	}
}
