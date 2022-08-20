package logger

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level string

const (
	Debug  Level = "debug"
	Info   Level = "info"
	Warn   Level = "warn"
	Error  Level = "error"
	DPanic Level = "dpanic"
	Panic  Level = "panic"
	Fatal  Level = "fatal"
)

type Format string

const (
	Text Format = "text"
	Json Format = "json"
)

var (
	// DefaultLevel 默认等级
	DefaultLevel = Info

	// DefaultFormat 默认输出格式
	DefaultFormat = Json

	// DefaultWriter 默认输出 Writer
	DefaultWriter = os.Stdout

	// DefaultCallerSkip 默认跳过 caller 的层级
	DefaultCallerSkip = 0
)

type Logger struct {
	Level      Level
	Format     Format
	Writer     io.Writer
	CallerSkip int
}

type Option func(*Logger)

func WithLevel(level Level) Option {
	return func(l *Logger) {
		l.Level = level
	}
}

func WithFormat(format Format) Option {
	return func(l *Logger) {
		l.Format = format
	}
}

func WithWriter(writer io.Writer) Option {
	return func(l *Logger) {
		l.Writer = writer
	}
}

func WithCallerSkip(skip int) Option {
	return func(l *Logger) {
		l.CallerSkip = skip
	}
}

func New(options ...Option) *zap.Logger {
	logger := Logger{
		Level:      DefaultLevel,
		Format:     DefaultFormat,
		Writer:     DefaultWriter,
		CallerSkip: DefaultCallerSkip,
	}

	for _, option := range options {
		option(&logger)
	}

	var (
		encoder zapcore.Encoder
	)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	switch logger.Format {
	case Text:
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	case Json:
		fallthrough
	default:
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	writeSyncer := zapcore.AddSync(logger.Writer)

	core := zapcore.NewCore(
		encoder,
		writeSyncer,
		logger.Level.Convert(),
	)

	zapLogger := zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(logger.CallerSkip),
	)

	return zapLogger
}

func (l Level) Convert() zapcore.Level {
	switch l {
	case Debug:
		return zapcore.DebugLevel
	case Info:
		return zapcore.InfoLevel
	case Warn:
		return zapcore.WarnLevel
	case Error:
		return zapcore.ErrorLevel
	case DPanic:
		return zapcore.DPanicLevel
	case Panic:
		return zapcore.PanicLevel
	case Fatal:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}
