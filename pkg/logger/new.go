package logger

import (
	"go-net-service/pkg/config"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func init() {
	logger = NewLogger().Sugar()
	_ = logger.Sync()
}

func NewLogger() *zap.Logger {
	core := newCore(getFilePath(), getLevel(), true)
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

func getLevel() zapcore.Level {
	if level, ok := levelMap[getLogLevel()]; ok {
		return level
	}
	return zapcore.DebugLevel
}

func getLogFilename() string {
	return strings.ToLower(config.GetString("app.name"))
}

func getLogLevel() string {
	return config.GetString("logger.level")
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
