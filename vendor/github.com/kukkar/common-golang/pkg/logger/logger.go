package logger

import (
	"strings"

	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var loggerImpls map[string]LogInterface

//
// InitiateLogger represents logger service which use zap library to print logs
//
func (this Config) InitiateLogger() error {

	loggerImpls = make(map[string]LogInterface)
	switch this.LoggerService {
	case "file":
	default:
		var zConfig zapConfig
		if this.OutputPaths == "" {
			this.OutputPaths = "stdout"
		}
		copier.Copy(&zConfig, &this)
		zLogger, err := getZapLogger(zConfig)
		if err != nil {
			return err
		}
		loggerImpls["zap"] = zLogger
	}
	return nil
}

func getZapLogger(config zapConfig) (LogInterface, error) {

	level := zapcore.Level(config.Level)
	outputPaths := strings.Split(config.OutputPaths, ",")
	errorPaths := strings.Split(config.ErrorOutputPaths, ",")
	var messageKey, levelKey, timeKey string // callerKey
	if config.EncoderConfig == (EncoderConfig{}) {
		messageKey = "message"
		levelKey = "level"
		timeKey = "time"
		//callerKey = "caller"
	} else {
		messageKey = config.EncoderConfig.MessageKey
		levelKey = config.EncoderConfig.LevelKey
		timeKey = config.EncoderConfig.TimeKey
		//callerKey = config.EncoderConfig.CallerKey
	}
	zConfig := zap.Config{
		Level:            zap.NewAtomicLevelAt(level),
		Development:      config.DevelopmentEnv,
		Encoding:         config.Encoding,
		OutputPaths:      outputPaths,
		ErrorOutputPaths: errorPaths,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   messageKey,
			LevelKey:     levelKey,
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			TimeKey:      timeKey,
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	var err error
	Logger, err = zConfig.Build()
	if err != nil {
		return nil, err
	}
	defer Logger.Sync()

	return &zapImpl{
		zap: Logger,
	}, nil
}
