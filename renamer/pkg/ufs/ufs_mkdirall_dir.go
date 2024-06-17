package ufs

import "os"

func MkdirAll(dir string) (err error) {
	return os.MkdirAll(dir, 0777)
}
