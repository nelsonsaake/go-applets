package main

import "fmt"

var (
	pubURL string = "ws://pubsubps.herokuapp.com/pub?id=foobar"
	subURL string = "ws://pubsubps.herokuapp.com/sub?id=foobar"
)

func getWSURL(ct ClientType) string {
	switch ct {
	case pub:
		return pubURL
	case sub:
		return subURL
	default:
		fmt.Println("Error selecting url, bad client type.")
	}
	return ""
}
