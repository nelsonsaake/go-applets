package sysout

import (
	"bytes"
	"fmt"
	"strings"
)

func _hasf(vs ...interface{}) bool {
	if len(vs) > 0 {
		if str, ok := vs[0].(string); ok {
			return strings.Contains(str, "%")
		}
	}
	return false
}

func _lnbrk(vs ...interface{}) string {
	if len(vs) > 0 {
		return "\n"
	}
	return ""
}

func _strvsf(f string, vs ...interface{}) string {
	vvs := make([]interface{}, 0)
	for _, v := range vs {
		vvs = append(vvs, fmtv(v))
	}
	return fmt.Sprintf(f, vvs...)
}

func _strvs(vs ...interface{}) string {
	b := bytes.Buffer{}
	write := b.WriteString
	for _, v := range vs {
		write(fmt.Sprintln(fmtv(v)))
	}
	return b.String()
}

func _str(msgs ...interface{}) string {
	b := bytes.Buffer{}
	write := b.WriteString

	write(_lnbrk(msgs...))
	defer write(_lnbrk(msgs...))

	switch {
	case _hasf(msgs...):
		write(_strvsf(msgs[0].(string), msgs[1:]...))
	default:
		write(_strvs(msgs...))
	}

	return b.String()
}

func tostring(msgs ...interface{}) string {
	return _str(msgs...)
}
