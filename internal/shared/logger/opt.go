package logger

import (
	"os"
)

type Option func(*defaultLogger) error

func OptNoop() Option {
	return func(logger *defaultLogger) error {
		logger.noopLogger = true
		return nil
	}
}

func MaskEnabled() Option {
	return func(logger *defaultLogger) error {
		logger.maskEnabled = true
		return nil
	}
}

func WithStdout() Option {
	return func(logger *defaultLogger) error {
		// Wire STD output for both type
		logger.writers = append(logger.writers, os.Stdout)
		return nil
	}
}

// WithLevel set level of logger
func WithLevel(level Level) Option {
	return func(logger *defaultLogger) error {
		logger.level = level
		return nil
	}
}
