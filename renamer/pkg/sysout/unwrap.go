package sysout

import (
	"strings"
)

func unwrap(txt, wrappers string) string {
	txt = strings.TrimSpace(txt)
	txt = strings.TrimPrefix(txt, wrappers)
	txt = strings.TrimSuffix(txt, wrappers)
	return txt
}
