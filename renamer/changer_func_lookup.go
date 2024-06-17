package main

import changecase "github.com/ku/go-change-case"

var changerFuncLookup map[string]func(string) string = map[string]func(string) string{
	"titleCase": changecase.Title,
	"camelCase": changecase.Camel,
	"index":     Index,
}
