package main

import "os"

func delFile(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		panic(err)
	}
}
