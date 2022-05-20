package logger

import (
	"fmt"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	helpers "github.com/hellogo/internal/common"
	"github.com/hellogo/internal/config"
)

type Level string

const (
	defaultPath     = "logs"
	defaultFileName = "hellogo.log"
)

const (
	LevelDebug Level = "debug"
	LevelInfo  Level = "info"
	LevelWarn  Level = "warn"
	LevelError Level = "error"
)

var LevelMap map[Level]int = map[Level]int{
	LevelDebug: 1 << 0,
	LevelInfo:  1 << 1,
	LevelWarn:  1 << 2,
	LevelError: 1 << 3,
}

/**
 * {@code zap.SugaredLogger} 实例
 */
var logger *zap.SugaredLogger

var _level Level = LevelInfo

func New(loggerConfig config.LoggerConfig) error {

	logLevels := mapLoggerLevels()

	writeSyncer, err := populateLoggerWriter(loggerConfig)
	if err != nil {
		return fmt.Errorf("parse the logger config error:%w", err)
	}
	encoder := populateLoggerEncoder(loggerConfig)
	level, ok := logLevels[loggerConfig.Level]
	if !ok {
		level = logLevels["info"]
	}
	_level = Level(loggerConfig.Level)

	core := zapcore.NewCore(encoder, writeSyncer, level)
	// _logger := zap.New(core)
	_logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	logger = _logger.Sugar()

	return nil
}

// LoggerFunc Logger instance
func LoggerFunc() *zap.SugaredLogger {
	return logger
}

func LoggerEnabled() bool {
	return nil != logger
}

// IsDebugEnabled 支持 debug 级别日志打印
func IsDebugEnabled() bool {
	level, ok := LevelMap[_level]
	if ok {
		return level <= LevelMap[LevelDebug]
	}

	return false
}

// IsInfoEnabled 支持 info 级别日志打印
func IsInfoEnabled() bool {
	level, ok := LevelMap[_level]
	if ok {
		return level <= LevelMap[LevelInfo]
	}

	return false
}

// IsWarnEnabled 支持 warn 级别日志打印
func IsWarnEnabled() bool {
	level, ok := LevelMap[_level]
	if ok {
		return level <= LevelMap[LevelWarn]
	}

	return false
}

// IsErrorEnabled 支持 error 级别日志打印
func IsErrorEnabled() bool {
	level, ok := LevelMap[_level]
	if ok {
		return level <= LevelMap[LevelError]
	}

	return false
}

func mapLoggerLevels() map[string]zapcore.Level {
	return map[string]zapcore.Level{
		"debug": zapcore.DebugLevel,
		"info":  zapcore.InfoLevel,
		"warn":  zapcore.WarnLevel,
		"error": zapcore.ErrorLevel,
	}
}

func populateLoggerEncoder(loggerConfig config.LoggerConfig) zapcore.Encoder {
	conf := zap.NewProductionEncoderConfig()
	conf.EncodeTime = zapcore.ISO8601TimeEncoder
	conf.EncodeLevel = zapcore.CapitalLevelEncoder

	return zapcore.NewJSONEncoder(conf)
}

func populateLoggerWriter(conf config.LoggerConfig) (zapcore.WriteSyncer, error) {
	conf.Path = helpers.Match("" == conf.Path, conf.Path, defaultPath)
	conf.FileName = helpers.Match("" == conf.FileName, conf.FileName, defaultFileName)

	if exist := helpers.Exists(conf.Path); !exist {
		if err := os.MkdirAll(conf.Path, os.ModePerm); err != nil {
			return nil, err
		}
	}

	_logger := populateLumberjackLogger(conf)

	syncer, err, done := handleMultiWriteSyncer(conf, _logger)
	if done {
		return syncer, err
	}

	return zapcore.AddSync(_logger), nil
}

func populateLumberjackLogger(conf config.LoggerConfig) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   filepath.Join(conf.Path, conf.FileName),
		MaxSize:    conf.MaxSize,
		MaxBackups: 0,
		MaxAge:     conf.MaxAgeDay,
		Compress:   conf.CompressEnabled,
	}
}

func handleMultiWriteSyncer(conf config.LoggerConfig, _logger *lumberjack.Logger) (zapcore.WriteSyncer, error, bool) {
	if conf.StdoutEnabled {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(_logger), zapcore.AddSync(os.Stdout)), nil, true
	}
	return nil, nil, false
}

// -----------------------------------------

func Logger() *zap.SugaredLogger {
	return logger
}

// ---------------------------------------

func Debug(message string, args ...any) {
	if LoggerEnabled() {
		logger.Debugf(message, args)
	}
}

func Info(message string, args ...any) {
	if LoggerEnabled() {
		logger.Infof(message, args)
	}
}

func Infow(message string, keysAndValues ...any) {
	if LoggerEnabled() {
		logger.Infow(message, keysAndValues)
	}
}

func Warn(message string, args ...any) {
	if LoggerEnabled() {
		logger.Warnf(message, args)
	}
}

func Warnw(message string, keysAndValues ...any) {
	if LoggerEnabled() {
		logger.Warnw(message, keysAndValues)
	}
}

func Error(message string, args ...any) {
	if LoggerEnabled() {
		logger.Errorf(message, args)
	}
}

func Errorw(message string, keysAndValues ...any) {
	if LoggerEnabled() {
		logger.Errorw(message, keysAndValues)
	}
}

func Fatal(message string, args ...any) {
	if LoggerEnabled() {
		logger.Fatalf(message, args)
	}
}

func Panic(message string, args ...any) {
	if LoggerEnabled() {
		logger.Panicf(message, args)
	}
}

func Sync() {
	_ = logger.Sync()
}
