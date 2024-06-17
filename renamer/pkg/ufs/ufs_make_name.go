package ufs

import (
	"fmt"
	"path/filepath"
	"time"
)

func MakeName(fileName string) string {

	extension := filepath.Ext(fileName)

	fileName = fmt.Sprint(time.Now().UnixNano())

	fileName += extension

	return fileName
}
