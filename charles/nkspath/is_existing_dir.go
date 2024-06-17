package nkspath

import (
	"errors"
	"os"
)

// is it a directory
// does the directory exist
// true, if both of the statement is true, otherwise false
func IsExistingDir(path string) bool {
	if info, err := os.Stat(path); err == nil {
		return true && info.IsDir()
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		panic(err)
	}
}
