package logger

import (
	"github.com/aclgo/grpc-jwt/config"
	"go.uber.org/zap"
)

type Logger interface {
	InitLogger()
	Debug()
	Debugf()
	Info()
	Infof(args ...any)
	Warn()
	Warnf()
	Error()
	Errorf()
	Fatal()
	Fatalf()
	Panic()
	Panicf()
}

type logger struct {
	config        *config.Config
	sugaredLogger zap.SugaredLogger
}

func NewLogger(cfg *config.Config) *logger {
	return &logger{
		config: cfg,
	}
}

func (l *logger) InitLogger() {
	// logLevel := zapcore.InfoLevel

	// logWriter := zapcore.AddSync(os.Stderr)

	// var encoderConfig zapcore.EncoderConfig

}

func (l *logger) Info(args ...any) {

}
func (l *logger) Infof() {

}
func (l *logger) Warn() {

}
func (l *logger) Warnf() {

}
func (l *logger) Fatal() {

}
func (l *logger) Fatalf() {

}

func (l *logger) Panic() {

}
func (l *logger) Panicf() {

}
