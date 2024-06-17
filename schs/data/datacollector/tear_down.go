package main

var teardownfuncs = []func(){}

func OnTearDown(f func()) {
	teardownfuncs = append(teardownfuncs, f)
}

func TearDown() {
	for _, f := range teardownfuncs {
		f()
	}
}
