package logx

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
)

type FileLoggerConfig struct {
	FilePath         string
	UseLocalTime     bool
	FileMaxSizeInMB  int
	FileMaxAgeInDays int
}

// Default Config Constants
const (
	defaultFilePath        = "logs/logs.json"
	defaultUseLocalTime    = false
	defaultFileMaxSizeInMB = 10
	defaultFileAgeInDays   = 30
)

// DefaultFileLoggerConfig returns a config with sensible defaults
func DefaultFileLoggerConfig() FileLoggerConfig {
	return FileLoggerConfig{
		FilePath:         defaultFilePath,
		UseLocalTime:     defaultUseLocalTime,
		FileMaxSizeInMB:  defaultFileMaxSizeInMB,
		FileMaxAgeInDays: defaultFileAgeInDays,
	}
}

func FileLogger(cfg FileLoggerConfig) io.Writer {
	return &lumberjack.Logger{
		Filename:  cfg.FilePath,
		LocalTime: cfg.UseLocalTime,
		MaxSize:   cfg.FileMaxSizeInMB,
		MaxAge:    cfg.FileMaxAgeInDays,
	}
}
