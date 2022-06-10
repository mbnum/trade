package trade

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func (e *Error) Error() string {
	if e != nil {
		return e.Out
	}
	return ""
}

func NewE(class int, out string, v ...interface{}) *Error {
	return &Error{Class: int32(class), Out: out, Stack: "created:" + strconv.Itoa(int(time.Now().UnixNano()/1e6)) + "," + fmt.Sprintf(strings.Repeat("%v,", len(v)), v...) + getMinifiedStack()}
}

func NewError(err error, v ...interface{}) *Error {
	pberr, ok := err.(*Error)
	if !ok {
		return &Error{Class: 500, Out: err.Error(), Stack: "created:" + strconv.Itoa(int(time.Now().UnixNano()/1e6)) + "," + fmt.Sprintf(strings.Repeat("%v,", len(v)), v...) + getMinifiedStack()}
	}
	pberr.Stack += "created:" + strconv.Itoa(int(time.Now().UnixNano()/1e6)) + "," + fmt.Sprintf(strings.Repeat("%v,", len(v)), v...) + getMinifiedStack()
	return pberr
}

func getMinifiedStack() string {
	stack := ""
	for i := 2; i < 5; i++ {
		_, fn, line, _ := runtime.Caller(i)
		if fn == "" {
			break
		}
		var split = strings.Split(fn, string(os.PathSeparator))
		var n int

		if len(split) >= 2 {
			n = len(split) - 2
		} else {
			n = len(split)
		}
		fn = strings.Join(split[n:], string(os.PathSeparator))

		stack += fmt.Sprintf("%s:%d", fn, line) + ","
	}
	return strings.TrimRight(stack, ",")
}
