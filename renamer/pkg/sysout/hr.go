package sysout

import "fmt"

func Hr() {
	hr := func(w int) (hr string) {
		for i := 0; i < w; i++ {
			hr += "-"
		}
		return
	}
	fmt.Println(hr(30))
	fmt.Println()
}
