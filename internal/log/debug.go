package log

import (
	"fmt"
	"os"
)

var (
	debugEnabled bool = false
)

func DebugEnabled() bool {
	return debugEnabled
}

func SetDebugEnabled(enabled bool) {
	debugEnabled = enabled
}

func Debugf(f string, a ...any) {
	if debugEnabled {
		fmt.Fprintf(os.Stderr, f, a...)
	}
}

func Debug(a ...any) {
	if debugEnabled {
		fmt.Fprintln(os.Stderr, a...)
	}
}

func Fatalf(f string, a ...any) {
	fmt.Fprintf(os.Stderr, f, a...)
	os.Exit(1)
}

func Fatal(a ...any) {
	fmt.Fprintln(os.Stderr, a...)
	os.Exit(1)
}
