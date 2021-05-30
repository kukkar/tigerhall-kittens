package logger

import (
	"errors"
	"fmt"
	"runtime"
	"strings"

	"github.com/kukkar/common-golang/pkg/logger/message"
	"github.com/kukkar/common-golang/pkg/utils"
)

func InfoSpecific(logType string, a ...interface{}) {

	loggerHandle, err := GetLoggerHandle(logType)
	if err != nil {
		fmt.Println("Skipping Log Info Message. " + err.Error())
		return
	}
	msg := Convert(a...)
	loggerHandle.Info(msg)
}

func Info(a ...interface{}) {
	InfoSpecific(defaultLoggerType, a...)
}

func DebugSpecific(logType string, a ...interface{}) {

	loggerHandle, err := GetLoggerHandle(logType)
	if err != nil {
		fmt.Println("Skipping Log Info Message. " + err.Error())
		return
	}
	msg := Convert(a...)
	loggerHandle.Info(msg)
}

func Debug(a ...interface{}) {
	DebugSpecific(defaultLoggerType, a...)
}

func ErrorSpecific(logType string, a ...interface{}) {

	loggerHandle, err := GetLoggerHandle(logType)
	if err != nil {
		fmt.Println("Skipping Log Info Message. " + err.Error())
		return
	}
	msg := Convert(a...)
	loggerHandle.Info(msg)
}

func Error(a ...interface{}) {
	ErrorSpecific(defaultLoggerType, a...)
}

func WarningSpecific(logType string, a ...interface{}) {

	loggerHandle, err := GetLoggerHandle(logType)
	if err != nil {
		fmt.Println("Skipping Log Info Message. " + err.Error())
		return
	}
	msg := Convert(a...)
	loggerHandle.Info(msg)
}

func Warning(a ...interface{}) {
	WarningSpecific(defaultLoggerType, a...)
}

// Convert converts application log to LogMsg format suitable for Logger
func Convert(a ...interface{}) message.LogMsg {
	paramLength := len(a)
	if paramLength == 0 {
		return message.LogMsg{
			Message: "Empty log param",
		}
	}
	if paramLength == 1 {
		//Only Log Message string is passed
		return message.LogMsg{
			Message: fmt.Sprintf("%s", a[0]),
		}
	}

	//First param is message string; Second param is request context
	vMsg, msgOk := a[0].(string)
	vRc, rcOk := a[1].(utils.RequestContext)

	if !msgOk || !rcOk {

		return message.LogMsg{
			Message: fmt.Sprintf("Erorr in parsing logging params for %v", a),
		}
	}
	return message.LogMsg{
		Message:       vMsg,
		TransactionID: vRc.TransactionID,
		SessionID:     vRc.SessionID,
		RequestID:     vRc.RequestID,
		AppID:         vRc.ClientAppID,
		UserID:        vRc.UserID,
		URI:           fmt.Sprintf("%s %s", strings.ToUpper(vRc.Method), vRc.URI),
		IP:            vRc.IP,
	}
}

func GetLoggerHandle(logType string) (LogInterface, error) {
	loggerHandle, ok := loggerImpls[logType]
	if !ok {
		return nil, errors.New("Undefined log type requested " + logType)
	}
	return loggerHandle, nil
}

//getStackTrace gets the stack trace for a called function.
func getStackTrace() []string {
	var sf []string
	j := 0
	for i := Skip; ; i++ {
		_, filePath, lineNumber, ok := runtime.Caller(i)
		if !ok || j >= CallingDepth {
			break
		}
		sf = append(sf, fmt.Sprintf("%s(%d)", filePath, lineNumber))
		j++
	}
	return sf
}
