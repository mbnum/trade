package trade

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

// Error returns string representation of an Error
func (e *Error) Error() string {
	if e == nil {
		return ""
	}

	b, _ := json.Marshal(e)
	return "#ERR " + string(b)
}

func NewError(class int, code E, v ...interface{}) *Error {
	var format, message string
	if len(v) == 0 {
		format = ""
	} else {
		var ok bool
		format, ok = v[0].(string)
		if !ok {
			format = strings.Repeat("%v", len(v))
		} else {
			v = v[1:]
		}
	}
	message = fmt.Sprintf(format, v...)
	message += " #STACK " + getMinifiedStack()

	e := &Error{}
	e.Class = int32(class)
	e.Debug = message
	e.Created = time.Now().UnixNano() / 1e6
	e.Code = code.String()
	return e
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
