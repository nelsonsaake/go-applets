package cout

import (
	"fmt"
)

func tostring(v interface{}) string {
	switch v.(type) {
	case error, string, int:
		return fmt.Sprint(v)
	}
	if s, ok := v.(fmt.Stringer); ok {
		return s.String()
	}
	return tojson(v)
}
