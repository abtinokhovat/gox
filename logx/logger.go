package logx

import (
	"context"
	"io"
	"log/slog"
	"os"
	"sync"
)

// Logger struct manages a singleton structured logger.
type Logger struct {
	mu      sync.Mutex
	writers []io.Writer
	logger  *slog.Logger
	options slog.HandlerOptions
}

type Config struct {
	Level slog.Level `koanf:"level"`
}

// Singleton instance
var (
	defaultLogger *Logger
	defaultOpts   = Config{Level: slog.LevelDebug}
)

func Default() *Logger {
	return defaultLogger
}

// init initializes the singleton logger with stdout.
func init() {
	defaultLogger = New(defaultOpts)
}

// New creates a logger with stdout as the default writer.
func New(cfg Config) *Logger {
	l := &Logger{
		writers: []io.Writer{os.Stdout},
		options: slog.HandlerOptions{
			Level: cfg.Level,
		},
	}

	l.recreate()
	return l
}

func (l *Logger) WithConfig(cfg Config) *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.options = slog.HandlerOptions{
		Level: cfg.Level,
	}
	l.recreate()
	return l
}

// WithFileLogger adds a file writer to the logger.
func (l *Logger) WithFileLogger(cfg FileLoggerConfig) *Logger {
	return l.WithWriters(FileLogger(cfg))
}

// WithWriters adds a new writer to the singleton logger.
func (l *Logger) WithWriters(w ...io.Writer) *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.writers = append(l.writers, w...)
	l.recreate()

	return l
}

// recreate updates the slog.Logger with all writers.
func (l *Logger) recreate() {
	multiWriter := io.MultiWriter(l.writers...)
	l.logger = slog.New(slog.NewJSONHandler(multiWriter, &l.options))
}

// Log returns the internal slog.Logger instance.
func (l *Logger) Log() *slog.Logger {
	return l.logger
}

func (l *Logger) log(ctx context.Context, level slog.Level, msg string, args ...any) {
	l.Log().Log(ctx, level, msg, args...)
}

// Debug calls [Logger.Debug] on the default logger.
func Debug(msg string, args ...any) {
	Default().log(context.Background(), slog.LevelDebug, msg, args...)
}

// Info calls [Logger.Info] on the default logger.
func Info(msg string, args ...any) {
	Default().log(context.Background(), slog.LevelInfo, msg, args...)
}

// Warn calls [Logger.Warn] on the default logger.
func Warn(msg string, args ...any) {
	Default().log(context.Background(), slog.LevelWarn, msg, args...)
}

// Error calls [Logger.Error] on the default logger.
func Error(msg string, args ...any) {
	Default().log(context.Background(), slog.LevelError, msg, args...)
}

// Log calls [Logger.Log] on the default logger.
func Log(ctx context.Context, level slog.Level, msg string, args ...any) {
	Default().log(ctx, level, msg, args...)
}
