package gfx

import "fmt"

// Log to standard output.
func Log(format string, a ...interface{}) {
	fmt.Printf(format+"\n", a...)
}

// Dump all of the arguments to standard output.
func Dump(a ...interface{}) {
	for _, v := range a {
		Log("%+v", v)
	}
}

// Printf formats according to a format specifier and writes to standard output.
func Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(format, a...)
}

// Sprintf formats according to a format specifier and returns the resulting string.
func Sprintf(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}
