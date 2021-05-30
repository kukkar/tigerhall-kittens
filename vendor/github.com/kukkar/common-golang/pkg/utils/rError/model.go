package rError

import "fmt"

// Error model consists of
// code i.e unique code either entered manually or table primary key
// msg i.e. actuall message need to display to user
// debug_msg i.e. the message added by developer to debugging. (optional)
// stack_trace i.e. the developer can get on demand error stack trace for debugging
// info is additional value need to be send with error
type Error struct {
	Code       int                    `json:"code" sql:"code" gorm:"column:code"`
	Msg        string                 `json:"msg" sql:"msg" gorm:"column:msg"`
	DebugMsg   string                 `json:"debug_msg,omitempty" sql:"-"`
	StackTrace []string               `json:"stack_trace,omitempty" sql:"-"`
	Info       map[string]interface{} `json:"info,omitempty" sql:"-"`
}

//Error implementing the error interface
func (e *Error) Error() string {
	val := fmt.Sprintf("Error: \ncode: %v \nmsg: %v", e.Code, e.Msg)
	if e.DebugMsg != "" {
		val += fmt.Sprintf("\ndebug_msg: %v", e.DebugMsg)
	}
	if len(e.StackTrace) > 0 {
		val += fmt.Sprintf("\nstack_trace: %v", e.StackTrace)
	}
	if e.Info != nil && len(e.Info) > 0 {
		val += fmt.Sprintf("\ninfo: %v", e.Info)
	}
	return val
}

//StackTrace model
type StackTrace struct {
	Func string
	File string
	Line int
}
