package main

import (
	"bytes"
	"encoding/json"
)

func tojson(v interface{}) string {
	b := bytes.Buffer{}
	encoder := json.NewEncoder(&b)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "\t")
	encoder.Encode(v)
	s := b.String()
	return s
}
