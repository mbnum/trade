package trade

import (
	context "context"
	"encoding/base64"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
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

func CtxGrpc(ctx *Context) context.Context {
	data, err := proto.Marshal(ctx)
	if err != nil {
		panic(fmt.Sprintf("unable to marshal ctx, %v", ctx))
	}
	cred64 := base64.StdEncoding.EncodeToString(data)
	return metadata.NewOutgoingContext(context.Background(), metadata.Pairs("ctx", cred64))
}

func GrpcToCtx(c context.Context) *Context {
	md, ok := metadata.FromIncomingContext(c)
	if !ok {
		md, ok = metadata.FromOutgoingContext(c)
		if !ok {
			return nil
		}
	}
	cred64 := strings.Join(md["ctx"], "")
	if cred64 == "" {
		return nil
	}
	data, err := base64.StdEncoding.DecodeString(cred64)
	if err != nil {
		panic(fmt.Sprintf("%v, %s: %s", err, "wrong base64 ", cred64))
	}

	ctx := &Context{}
	if err = proto.Unmarshal(data, ctx); err != nil {
		panic(fmt.Sprintf("%v, %s: %s", err, "unable to unmarshal ctx ", cred64))
	}
	return ctx
}
