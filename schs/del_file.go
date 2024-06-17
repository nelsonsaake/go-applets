package main

import (
	"fmt"
	"projects/saelections/pkg/ufs"
)

func DelFile(file string) {
	if err := ufs.DelFile(file); err != nil {
		fmt.Println(err)
	}
}
