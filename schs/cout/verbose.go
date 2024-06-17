package cout

var _v bool

func SetVerbose(v bool) {
	_v = v
}

func IsVerbose() bool {
	return _v
}

func Verbose(v ...interface{}) {
	if _v {
		Print(v...)
	}
}
