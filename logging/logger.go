package logging

import (
	"io"
	"os"
	"path"
	"sync"

	"github.com/maestre3d/gtavd/global"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	DefaultLoggingFile = "gtavd.log"

	defaultMaxFileSize  = 128 // Megabytes
	defaultDayRetention = 14
)

var (
	DefaultLoggingPath = global.DefaultPath + "/logs"

	Logger     MultiLogger
	onceLogger = sync.Once{}
)

type MultiLogger struct {
	zerolog.Logger
}

func init() {
	onceLogger.Do(func() {
		tryCreateConfigPath()
		fileLogger := &lumberjack.Logger{
			Filename:   path.Join(DefaultLoggingPath, DefaultLoggingFile),
			MaxSize:    defaultMaxFileSize,
			MaxAge:     defaultDayRetention,
			MaxBackups: 0,
			LocalTime:  false,
			Compress:   true,
		}
		writers := io.MultiWriter(zerolog.ConsoleWriter{Out: os.Stderr}, fileLogger)
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		Logger = MultiLogger{
			Logger: zerolog.New(writers).With().Timestamp().Logger(),
		}
	})
}

func tryCreateConfigPath() {
	_ = os.MkdirAll(DefaultLoggingPath, 0750)
}

func Info() *zerolog.Event {
	return Logger.Info()
}

func Debug() *zerolog.Event {
	return Logger.Debug()
}

func Error() *zerolog.Event {
	return Logger.Error()
}
