package services

import (
	"fmt"
	"testing"
)

func TestGroupFiles(t *testing.T) {
	var testDir = `D:\downloads\Votee`

	for _, v := range GroupFiles(testDir, 10) {

		for _, vv := range v {
			fmt.Println(vv)
		}

		fmt.Println()
	}
}
