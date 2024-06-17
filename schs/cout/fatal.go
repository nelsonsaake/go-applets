package cout

import "os"

func Fatal(msgs ...interface{}) {
	Print(msgs...)
	os.Exit(-1)
}
