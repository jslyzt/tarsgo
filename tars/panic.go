package tars

import (
	"fmt"
	"os"

	"tarsgo/tars/util/debug"
	"tarsgo/tars/util/rogger"
)

//////////////////////////////////////////////////////////////////////////////
type PanicFunc func(interface{}) bool

var (
	FPanic PanicFunc
)

//////////////////////////////////////////////////////////////////////////////
// CheckPanic used to dump stack info to file when catch panic
func CheckPanic() {
	if r := recover(); r != nil {
		if FPanic != nil && FPanic(r) {
			return
		}
		var msg string
		if err, ok := r.(error); ok {
			msg = err.Error()
		} else {
			msg = fmt.Sprintf("%#v", r)
		}
		debug.DumpStack(true, "panic", msg)
		rogger.FlushLogger()
		os.Exit(-1)
	}
}
