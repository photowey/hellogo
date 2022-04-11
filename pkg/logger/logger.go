package logger

import (
	"fmt"
	"os"
	"path/filepath"

	helpers "github.com/hellogo/internal/common"
	"github.com/hellogo/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	defaultPath     = "logs"
	defaultFileName = "hellogo.log"
)

/**
 * {@code zap.SugaredLogger} 实例
 */
var logger *zap.SugaredLogger

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

	core := zapcore.NewCore(encoder, writeSyncer, level)
	_logger := zap.New(core)
	logger = _logger.Sugar()

	return nil
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

func Debug(message string, args ...interface{}) {
	logger.Debugf(message, args)
}

func Info(message string, args ...interface{}) {
	logger.Infof(message, args)
}

func Infow(message string, keysAndValues ...interface{}) {
	logger.Infow(message, keysAndValues)
}

func Warn(message string, args ...interface{}) {
	logger.Warnf(message, args)
}

func Warnw(message string, keysAndValues ...interface{}) {
	logger.Warnw(message, keysAndValues)
}

func Error(message string, args ...interface{}) {
	logger.Errorf(message, args)
}

func Errorw(message string, keysAndValues ...interface{}) {
	logger.Errorw(message, keysAndValues)
}

func Fatal(message string, args ...interface{}) {
	logger.Fatalf(message, args)
}

func Panic(message string, args ...interface{}) {
	logger.Panicf(message, args)
}
