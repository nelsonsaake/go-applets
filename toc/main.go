package main

import (
	_ "embed"
	"io/ioutil"
	"strings"
)

//go:embed outline
var outline string

//go:embed toc
var toc string

func lines(str string) []string {
	lines := []string{}
	lines = strings.Split(str, "\r\n")
	return lines
}

func main() {
	outlineLines := lines(outline)
	tocLines := lines(toc)

	i := 0
	for itoc, xtoc := range tocLines {

		if i < len(outlineLines) {
			if xtoc == outlineLines[i] {
				i++
				continue
			}
		}

		tocLines[itoc] = "\t" + xtoc
	}

	toc := strings.Join(tocLines, "\r\n")
	print(toc)

	ioutil.WriteFile("new toc.txt", []byte(toc), 0666)
}
