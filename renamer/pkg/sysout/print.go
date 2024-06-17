// sysout : pretty prints, coverts structs to json before prints them to out

package sysout

func Print(msgs ...interface{}) {
	print(tostring(msgs...))
}
