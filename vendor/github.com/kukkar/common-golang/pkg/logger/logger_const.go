package logger

const (
	DEBUG_LEVEL   int8 = 4
	INFO_LEVEL    int8 = 3
	TRACELEVEL    int8 = 2
	WARNING_LEVEL int8 = 1
	ERROR_LEVEL   int8 = 0

	//Define color specific constants.
	greenColor        = "\x1b[32m"
	redColor          = "\x1b[91m"
	yellowColor       = "\x1b[33m"
	blueColor         = "\x1b[34m"
	pinkColor         = "\x1b[91m"
	purpleColor       = "\x1b[35m"
	lightBlueColor    = "\x1b[36m"
	defaultStyle      = "\x1b[0m"
	lightGrayColor    = "\x1b[30m"
	defaultLoggerType = "zap"
	Skip              = 4
	CallingDepth      = 5
)

var LOG_LEVEL_MAP = map[int]string{
	5: "debug",
	4: "info",
	1: "error",
}

var logFields = []string{"URI", "RequestID", "IP", "TimeTaken", "Caller"}
