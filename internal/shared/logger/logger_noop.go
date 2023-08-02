package logger

import "context"

type NoopContextLogger struct{}

// Debug ...
func (n *NoopContextLogger) Debug(context.Context, string, ...Field) {}

// Info ...
func (n *NoopContextLogger) Info(context.Context, string, ...Field) {}

// Warn ...
func (n *NoopContextLogger) Warn(context.Context, string, ...Field) {}

// Error ...
func (n *NoopContextLogger) Error(context.Context, string, ...Field) {}

// Fatal ...
func (n *NoopContextLogger) Fatal(context.Context, string, ...Field) {}

// Panic ...
func (n *NoopContextLogger) Panic(context.Context, string, ...Field) {}

// Close ...
func (n *NoopContextLogger) Close() error {
	return nil
}
