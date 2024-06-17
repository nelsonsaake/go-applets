package cout

import "fmt"

func endln(vs []interface{}) {
	if len(vs) != 0 {
		fmt.Println()
	}
}
