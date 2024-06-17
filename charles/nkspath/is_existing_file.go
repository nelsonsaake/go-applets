package nkspath

import (
	"errors"
	"os"
)

// is it a file
// does the file exist
// true, if both of the statement is true, otherwise false
func IsExistingFile(path string) bool {
	if info, err := os.Stat(path); err == nil {
		return true && !info.IsDir()
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		panic(err)
	}
}
