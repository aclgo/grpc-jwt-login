package logger

import (
	"os"

	"github.com/aclgo/grpc-jwt/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	InitLogger()
	Debug(args ...any)
	Debugf(template string, args ...any)
	Info(args ...any)
	Infof(template string, args ...any)
	Warn(args ...any)
	Warnf(template string, args ...any)
	Error(args ...any)
	Errorf(template string, args ...any)
	Fatal(args ...any)
	Fatalf(template string, args ...any)
	DPanic(args ...any)
	DPanicf(template string, args ...any)
}

type logger struct {
	config        *config.Config
	sugaredLogger *zap.SugaredLogger
}

func NewapiLogger(cfg *config.Config) *logger {
	return &logger{
		config: cfg,
	}
}

var mapLevel = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func (l *logger) getLogLevel() zapcore.Level {
	level, exist := mapLevel[l.config.LogLevel]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

func (l *logger) InitLogger() {
	logLevel := l.getLogLevel()

	logWriter := zapcore.AddSync(os.Stderr)

	var encConfig zapcore.EncoderConfig

	if l.config.ServerMode == "dev" {
		encConfig = zap.NewDevelopmentEncoderConfig()
	} else {
		encConfig = zap.NewProductionEncoderConfig()
	}

	encConfig.LevelKey = "LEVEL"
	encConfig.CallerKey = "CALLER"
	encConfig.TimeKey = "TIME"
	encConfig.NameKey = "NAME"
	encConfig.MessageKey = "MESSAGE"
	encConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var encoder zapcore.Encoder

	if l.config.LogEncoding == "console" {
		encoder = zapcore.NewConsoleEncoder(encConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encConfig)
	}

	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(logLevel))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	l.sugaredLogger = logger.Sugar()
	if err := l.sugaredLogger.Sync(); err != nil {
		l.sugaredLogger.Error(err)
	}
}

func (l *logger) Debug(args ...any) {
	l.sugaredLogger.Debug(args...)
}
func (l *logger) Debugf(template string, args ...any) {
	l.sugaredLogger.Debugf(template, args...)
}

func (l *logger) Info(args ...any) {
	l.sugaredLogger.Info(args...)
}

func (l *logger) Infof(template string, args ...any) {
	l.sugaredLogger.Infof(template, args...)
}
func (l *logger) Warn(args ...any) {
	l.sugaredLogger.Warn(args...)
}
func (l *logger) Warnf(template string, args ...any) {
	l.sugaredLogger.Warnf(template, args...)
}

func (l *logger) Error(args ...any) {
	l.sugaredLogger.Error(args...)
}
func (l *logger) Errorf(template string, args ...any) {
	l.sugaredLogger.Errorf(template, args...)
}
func (l *logger) Fatal(args ...any) {
	l.sugaredLogger.Fatal(args...)
}
func (l *logger) Fatalf(template string, args ...any) {
	l.sugaredLogger.Fatalf(template, args...)
}

func (l *logger) DPanic(args ...any) {
	l.sugaredLogger.DPanic(args...)
}
func (l *logger) DPanicf(template string, args ...any) {
	l.sugaredLogger.DPanicf(template, args...)
}
