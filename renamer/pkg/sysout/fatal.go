package sysout

import (
	"os"
	"testing"
)

func _tfatal(t *testing.T, msgs ...interface{}) {
	t.Fatal(tostring(msgs...))
}

func _dfatal(msgs ...interface{}) {
	print(tostring(msgs...))
	os.Exit(-1)
}

func Fatal(msgs ...interface{}) {
	switch {
	case len(msgs) > 0:
		switch msgs[0].(type) {
		case *testing.T:
			_tfatal(msgs[0].(*testing.T), msgs[1:]...)
		default:
			_dfatal(msgs...)
		}
	}
}
