package main

import "os"

func cleanFile(fileName string) {
	if err := os.Truncate(fileName, 0); err != nil {
		panic(err)
	}
}
