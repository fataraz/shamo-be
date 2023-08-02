package logger

import (
	"context"
	"fmt"
	"sync"
)

// Logger ...
type Logger interface {
	Debug(ctx context.Context, message string, fields ...Field)
	Info(ctx context.Context, message string, fields ...Field)
	Warn(ctx context.Context, message string, fields ...Field)
	Error(ctx context.Context, message string, fields ...Field)
	Fatal(ctx context.Context, message string, fields ...Field)
	Panic(ctx context.Context, message string, fields ...Field)
	Close() error
}

// Field ...
type Field struct {
	Key string
	Val interface{}
}

// ctxKeyLogger ...
type ctxKeyLogger struct{}

var ctxKey = ctxKeyLogger{}

// Context ...
type Context struct {
	ServiceName    string `json:"_app_name"`
	ServiceVersion string `json:"_app_version"`
	ServicePort    int    `json:"_app_port"`
	ThreadID       string `json:"_app_thread_id"`
	XRequestID     string `json:"_x_request_id"`
	XAgent         string `json:"_x_agent"`
	Tag            string `json:"_app_tag"`
	Request        Field  `json:"_request,omitempty"`
	Response       Field  `json:"_response,omitempty"`
	Error          string `json:"_error,omitempty"`

	ReqMethod string `json:"_app_method"`
	ReqURI    string `json:"_app_uri"`

	AdditionalData map[string]interface{} `json:"_app_data,omitempty"`
}

var instance Logger
var once sync.Once

// GetLogger ...
func GetLogger() Logger {
	once.Do(func() {
		instance = SetupLoggerFile()
	})
	return instance
}

// SetupLoggerFile ...
func SetupLoggerFile() Logger {
	fmt.Println("Try newLogger File...")

	var level Level
	var opt = make([]Option, 0)
	opt = append(opt, WithStdout())
	opt = append(opt, WithLevel(level))

	log, err := newLogger(opt...)
	if err != nil {
		panic(fmt.Errorf("init legacy logger with mode stdout error: %w", err))
	}

	return log
}

// InjectCtx ...
func InjectCtx(parent context.Context, ctx Context) context.Context {
	if parent == nil {
		return InjectCtx(context.Background(), ctx)
	}

	return context.WithValue(parent, ctxKey, ctx)
}

// ExtractCtx ...
func ExtractCtx(ctx context.Context) Context {
	if ctx == nil {
		return Context{}
	}

	val, ok := ctx.Value(ctxKey).(Context)
	if !ok {
		return Context{}
	}

	return val
}
