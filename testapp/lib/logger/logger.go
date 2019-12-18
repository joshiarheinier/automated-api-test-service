package logger

import (
	"context"
	"fmt"
	"github.com/joshia/automated-api-test-service/testapp/config"
	logcfg "github.com/joshia/automated-api-test-service/testapp/config/structs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"time"
)

var (
	logger *logrus.Logger
)

type Fields map[string]interface{}
func setOutputFile(filepath string, filename string, maxSize int, maxBackup int, maxAge int, compress bool) io.Writer {
	return &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", filepath, filename),
		MaxSize:    maxSize, // megabytes
		MaxBackups: maxBackup,
		MaxAge:     maxAge,   //days
		Compress:   compress, // disabled by default
	}
}

func init()  {
	logger = logrus.New()
	New(config.NewConfig().Logger)
}

func New(conf *logcfg.LogConfiguration) {
	fileFormatter := &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}
	stdoutFormatter := &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	errorLogFileWriter := setOutputFile(conf.FilePath, conf.ErrorLogFileName, conf.MaxSize, conf.MaxBackup, conf.MaxAge, conf.Compress)
	debugLogFileWriter := setOutputFile(conf.FilePath, conf.DebugLogFileName, conf.MaxSize, conf.MaxBackup, conf.MaxAge, conf.Compress)
	logFileWriter := setOutputFile(conf.FilePath, conf.LogFileName, conf.MaxSize, conf.MaxBackup, conf.MaxAge, conf.Compress)
	logger.AddHook(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.DebugLevel: debugLogFileWriter,
			logrus.InfoLevel:  logFileWriter,
			logrus.WarnLevel:  logFileWriter,
			logrus.ErrorLevel: errorLogFileWriter,
			logrus.FatalLevel: errorLogFileWriter,
		}, fileFormatter),
	)
	logger.SetFormatter(stdoutFormatter)
	logger.SetLevel(getLevel(conf.Level))
}
func getLevel(level string) logrus.Level {
	if level == "error" {
		return logrus.ErrorLevel
	} else if level == "info" {
		return logrus.InfoLevel
	} else if level == "debug" {
		return logrus.DebugLevel
	}
	return logrus.ErrorLevel
}
func Debug(args ...interface{}) {
	logger.Debug(args)
}
func Info(args ...interface{}) {
	logger.Info(args)
}
func Error(args ...interface{}) {
	logger.Error(args)
}
func Panic(args ...interface{}) {
	logger.Panic(args)
}
func WithContext(ctx context.Context) *logrus.Entry {
	return logger.WithContext(ctx)
}
func WithField(key string, value interface{}) *logrus.Entry {
	return logger.WithField(key, value)
}
func WithFields(fields Fields) *logrus.Entry {
	entry := logrus.NewEntry(logrus.StandardLogger())
	for k, v := range fields {
		entry = entry.WithField(k, v)
	}
	return entry
}
func WithTime(t time.Time) *logrus.Entry {
	return logger.WithTime(t)
}