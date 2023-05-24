package trade

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// TODO name to id

func (e *Error) Error() string {
	if e != nil {
		return e.Message
	}
	return ""
}

func (e *Error) GRPCStatus() *status.Status {
	s, _ := status.New(codes.Internal, e.Message).WithDetails(e)
	return s
}

func GrpcToErr(err error) *Error {
	if err == nil {
		return nil
	}

	se, ok := err.(interface{ GRPCStatus() *status.Status })
	if !ok {
		return &Error{Class: 500, Message: err.Error()}
	}
	s := se.GRPCStatus()

	details := s.Details()
	if len(details) == 0 {
		return &Error{Class: 500, Message: s.Message()}
	}

	e, ok := details[0].(*Error)
	if !ok {
		return &Error{Class: 500, Message: s.Message()}
	}
	return e
}

func NewE(class int, message string, v ...interface{}) *Error {
	e := &Error{Class: int32(class), Message: message}
	if message == "" && len(v) > 0 {
		if b, err := json.Marshal(v[0]); err == nil {
			e.Prob = string(b)
			v = v[1:]
		}
	}
	e.Stack = "created:" + strconv.Itoa(int(time.Now().UnixNano()/1e6)) + "," + fmt.Sprintf(strings.Repeat("%v,", len(v)), v...) + getMinifiedStack()
	return e
}

func NewError(err error, v ...interface{}) *Error {
	pberr, ok := err.(*Error)
	if !ok {
		return &Error{Class: 500, Message: err.Error(), Stack: "created:" + strconv.Itoa(int(time.Now().UnixNano()/1e6)) + "," + fmt.Sprintf(strings.Repeat("%v,", len(v)), v...) + getMinifiedStack()}
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
