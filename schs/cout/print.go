// cout : pretty prints, coverts structs to json before prints them to out

package cout

import (
	"fmt"
	"strings"
)

func _prinf(msgs ...interface{}) bool {
	if len(msgs) > 1 {
		if str, ok := msgs[0].(string); ok {
			if strings.Contains(str, "%") {
				js := make([]interface{}, 0)
				for _, v := range msgs {
					js = append(js, tostring(v))
				}
				fmt.Printf(str, js[1:]...)
				return true
			}
		}
	}
	return false
}

func _print(msgs ...interface{}) {
	for _, v := range msgs {
		switch v.(type) {
		case error:
			fmt.Println(v)
		default:
			fmt.Println(tostring(v))
		}
	}
}

func Print(msgs ...interface{}) {
	endln(msgs)
	defer endln(msgs)
	if !_prinf(msgs) {
		_print(msgs...)
	}
}
