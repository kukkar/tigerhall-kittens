package logger

type Config struct {
	LoggerService    string        `json:"LoggerService"`
	Level            int           `json:"Level"`
	DevelopmentEnv   bool          `json:"DevelopmentEnv"`
	Encoding         string        `json:"Encoding"`
	InitialField     string        `json:"InitialField"`
	EncoderConfig    EncoderConfig `json:"EncoderConfig"`
	OutputPaths      string        `json:"OutputPaths"`
	ErrorOutputPaths string        `json:"ErrorOutputPaths"`
	AppName          string        `json:"AppName"`
}

type EncoderConfig struct {
	MessageKey     string `json:"MessageKey"`
	LevelKey       string `json:"LevelKey"`
	NameKey        string `json:"NameKey"`
	TimeKey        string `json:"TimeKey"`
	CallerKey      string `json:"CallerKey"`
	StacktraceKey  string `json:"StacktraceKey"`
	CallStackKey   string `json:"CallStackKey"`
	ErrorKey       string `json:"ErrorKey"`
	EncodeTime     string `json:"TimeEncoder"`
	FileKey        string `json:"FileKey"`
	EncodeLevel    string `json:"LevelEncoder"`
	EncodeDuration string `json:"DurationEncoder"`
	EncodeCaller   string `json:"CallerEncoder"`
	EncodeName     string `json:"NameEncoder"`
}

type zapConfig struct {
	Level            int           `json:"Level"`
	DevelopmentEnv   bool          `json:"DevelopmentEnv"`
	Encoding         string        `json:"Encoding"`
	InitialField     string        `json:"InitialField"`
	EncoderConfig    EncoderConfig `json:"EncoderConfig"`
	OutputPaths      string        `json:"OutputPaths"`
	ErrorOutputPaths string        `json:"ErrorOutputPaths"`
	AppName          string        `json:"AppName"`
}

type zapEncoderConfig struct {
	MessageKey     string `json:"MessageKey"`
	LevelKey       string `json:"LevelKey"`
	NameKey        string `json:"NameKey"`
	TimeKey        string `json:"TimeKey"`
	CallerKey      string `json:"CallerKey"`
	StacktraceKey  string `json:"StacktraceKey"`
	CallStackKey   string `json:"CallStackKey"`
	ErrorKey       string `json:"ErrorKey"`
	EncodeTime     string `json:"TimeEncoder"`
	FileKey        string `json:"FileKey"`
	EncodeLevel    string `json:"LevelEncoder"`
	EncodeDuration string `json:"DurationEncoder"`
	EncodeCaller   string `json:"CallerEncoder"`
	EncodeName     string `json:"NameEncoder"`
}
