package rError

import (
	"context"
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

const (
	packageName        = "github.com/ralali/rl-go-microservice/sdk/rError"
	DebugKey           = "debug"
	StackTraceDepthKey = "trace_depth"
)

//newError will create new wrapped error
func newError(c context.Context, err error, code int, debugMsg ...string) *Error {

	var tracePath []string

	//if debug key is passed then only check stack trace
	if IsDebug(c) {
		limit := getTraceDepth(c)
		st := GetStackTrace(3, limit)
		for _, v := range st {
			curTrace := fmt.Sprintf("%v:%v %v()", v.File, v.Line, v.Func)
			tracePath = append(tracePath, curTrace)
		}
	}

	if err != nil {
		debugMsg = append(debugMsg, err.Error())
	}
	dMsg := strings.Join(debugMsg, " ")

	//setting msg from error map based on code
	msg, exists := errMap[code]
	if !exists {
		if err != nil {
			msg = err.Error()
		} else {
			msg = dMsg
		}
	}

	return &Error{
		Code:       code,
		Msg:        msg,
		DebugMsg:   dMsg,
		StackTrace: tracePath,
		Info:       make(map[string]interface{}),
	}
}

//IsDebug will check debug send in context
func IsDebug(c context.Context) (debug bool) {
	if c != nil {
		_, debug = c.Value(DebugKey).(string)
	}
	return
}

//getTraceDepth will return stack trace depth. Default is 3.
func getTraceDepth(c context.Context) (depth int) {
	if c != nil {
		dVal, _ := c.Value(StackTraceDepthKey).(string)
		depth, _ = strconv.Atoi(dVal)
	}
	if depth == 0 {
		depth = 3
	}
	return
}

//GetStackTrace : Get function name, file name and line no of the caller function
//Depth is the value from which it will start searching in the stack
func GetStackTrace(start int, limit int) (st []StackTrace) {
	counter := 0
	for i := start; ; i++ {
		if pc, file, line, ok := runtime.Caller(i); ok {
			if strings.Contains(file, packageName) {
				continue
			}

			fileName := strings.Split(file, "github.com")
			if len(fileName) > 1 {
				file = fileName[1]
			}
			_, funcName := packageFuncName(pc)

			var curST = StackTrace{Func: funcName, File: file, Line: line}
			st = append(st, curST)
			counter++
			if counter == limit {
				break
			}
		} else {
			break
		}
	}
	return
}

//packageFuncName : Package and function name from package counter
func packageFuncName(pc uintptr) (packageName string, funcName string) {
	if f := runtime.FuncForPC(pc); f != nil {
		funcName = f.Name()
		if ind := strings.LastIndex(funcName, "/"); ind > 0 {
			packageName += funcName[:ind+1]
			funcName = funcName[ind+1:]
		}
		if ind := strings.Index(funcName, "."); ind > 0 {
			packageName += funcName[:ind]
			funcName = funcName[ind+1:]
		}
	}
	return
}
