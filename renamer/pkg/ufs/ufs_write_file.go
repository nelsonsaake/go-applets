package ufs

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

func WriteFile(fpath, content string) (err error) {
	fdir := strings.TrimSuffix(fpath, filepath.Base(fpath))
	MkdirAll(fdir)
	err = ioutil.WriteFile(fpath, []byte(content), 7777)
	return
}
