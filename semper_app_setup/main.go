package main

import (
	_ "embed"
	"fmt"
	"projects/semper-server/pkg/syscmds"
	"strings"
)

//go:embed script.bat
var script string

var cmds []string

func init() {
	cmds = strings.Split(script, "\r\n")
}

func main() {
	for _, cmd := range cmds {
		if err := syscmds.Run(cmd); err != nil {
			fmt.Println(err)
		}
	}
}
