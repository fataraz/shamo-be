package logger

import (
	"reflect"
	"time"
)

const (
	// LogTypeSYS ...
	LogTypeSYS = "SYS"
)

// separator ...
const separator = "|"

var (
	TypeSliceOfBytes = reflect.TypeOf([]byte(nil))
	TypeTime         = reflect.TypeOf(time.Time{})
)
