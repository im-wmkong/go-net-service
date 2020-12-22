package logger

import (
	"go-net-service/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var logger *zap.SugaredLogger

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func Initialize() {
	filePath := getFilePath()
	level := getLoggerLevel()
	logger = NewLogger(filePath, level).Sugar()
	_ = logger.Sync()
}

func NewLogger(filePath string, level zapcore.Level) *zap.Logger {
	core := newCore(filePath, level, true)
	return zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.Development())
}

func newCore(filePath string, level zapcore.Level, compress bool) zapcore.Core {
	encoderConfig := getEncoderConfig()

	hook := lumberjack.Logger{
		Filename:   filePath, // 日志文件路径
		Compress:   compress, // 是否压缩
		MaxSize:    config.GetInt("logger.max_size"),
		MaxBackups: config.GetInt("logger.max_backup"),
		MaxAge:     config.GetInt("logger.max_age"),
	}

	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		level,
	)
}

func getFilePath() string {
	return "storage/logs/" + getLogFilename() + ".log"
}

func getLoggerLevel() zapcore.Level {
	if level, ok := levelMap[getLogLevel()]; ok {
		return level
	}
	return zapcore.DebugLevel
}

func getLogFilename() string {
	return config.GetString("app.name")
}

func getLogLevel() string {
	return config.GetString("app.log_level")
}

func getEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
		EncodeTime:     timeEncoder,
	}
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	logger.Debugw(msg, keysAndValues...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	logger.Infow(msg, keysAndValues...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	logger.Warnf(template, args...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	logger.Warnw(msg, keysAndValues...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	logger.Errorw(msg, keysAndValues...)
}

func Panic(args ...interface{}) {
	logger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	logger.Panicf(template, args...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	logger.Panicw(msg, keysAndValues...)
}
