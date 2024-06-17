package cout

import (
	"errors"
	"testing"
)

func TestPrintjson(t *testing.T) {
	Print(errors.New("an error"))
	Print("string ")
}
