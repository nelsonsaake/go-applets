package main

import (
	"fmt"
	"strings"
)

type ClientType int

const (
	pub ClientType = 1
	sub            = 2
	bad            = 3
)

//
func getCTStr() string {
	fmt.Print("Client:>> ")
	var ln string
	fmt.Scanln(&ln)

	ln = strings.ToLower(ln)
	ln = strings.Trim(ln, "\n")
	return ln
}

//
func getCT() ClientType {
	ln := getCTStr()

	switch ln {
	case "pub":
		return pub
	case "sub":
		return sub
	default:
		fmt.Print(ln)
		return bad
	}
}
