package logger

import (
	"github.com/kukkar/common-golang/pkg/logger/message"
)

var _ LogInterface = (*zapImpl)(nil)

type LogInterface interface {
	Trace(msg message.LogMsg)
	Warning(msg message.LogMsg)
	Info(msg message.LogMsg)
	Error(msg message.LogMsg)
	Debug(msg message.LogMsg)
	Profile(msg message.LogMsg)
}
