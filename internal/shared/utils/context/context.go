package context

import (
	"context"
	"errors"
	"github.com/spf13/cast"
	"shamo-be/internal/shared/models"
	"time"

	Map "github.com/orcaman/concurrent-map"
	Logger "shamo-be/internal/shared/logger"
)

// Context ...
type Context struct {
	Map                       Map.ConcurrentMap
	Logger                    Logger.Logger
	RequestTime               time.Time
	UserSession               models.AccountSession
	XRequestID, XAgent        string
	AppName, AppVersion, IP   string
	Port                      int
	SrcIP, URL, Method        string
	Header, Request, Response interface{}
	ErrorMessage              string
	ResponseCode              int
}

// New ...
func New(logger Logger.Logger) *Context {
	return &Context{
		RequestTime: time.Now(),
		Logger:      logger,
		Map:         Map.New(),
		Header:      map[string]interface{}{},
		Request:     struct{}{},
	}
}

// SetXRequestID ...
func (s *Context) SetXRequestID(xRequestID string) *Context {
	s.XRequestID = xRequestID
	return s
}

// SetXAgent ...
func (s *Context) SetXAgent(xAgent string) *Context {
	s.XAgent = xAgent
	return s
}

// SetMethod ...
func (s *Context) SetMethod(method string) *Context {
	s.Method = method
	return s
}

// SetAppName ...
func (s *Context) SetAppName(appName string) *Context {
	s.AppName = appName
	return s
}

// SetAppVersion ...
func (s *Context) SetAppVersion(appVersion string) *Context {
	s.AppVersion = appVersion
	return s
}

// SetURL ...
func (s *Context) SetURL(url string) *Context {
	s.URL = url
	return s
}

// SetIP ...
func (s *Context) SetIP(ip string) *Context {
	s.IP = ip
	return s
}

// SetPort ...
func (s *Context) SetPort(port int) *Context {
	s.Port = port
	return s
}

// SetSrcIP ...
func (s *Context) SetSrcIP(srcIp string) *Context {
	s.SrcIP = srcIp
	return s
}

// SetHeader ...
func (s *Context) SetHeader(header interface{}) *Context {
	s.Header = header
	return s
}

// SetRequest ...
func (s *Context) SetRequest(request interface{}) *Context {
	s.Request = request
	return s
}

// SetRequestTime ...
func (s *Context) SetRequestTime(request time.Time) *Context {
	s.RequestTime = request
	return s
}

// SetErrorMessage ...
func (s *Context) SetErrorMessage(errorMessage string) *Context {
	s.ErrorMessage = errorMessage
	return s
}

// SetResponseCode ...
func (s *Context) SetResponseCode(responseCode int) *Context {
	s.ResponseCode = responseCode
	return s
}

// Get ...
func (s *Context) Get(key string) (data interface{}, err error) {
	data, ok := s.Map.Get(key)
	if !ok {
		err = errors.New("not found")
	}
	return
}

// Put ...
func (s *Context) Put(key string, data interface{}) {
	s.Map.Set(key, data)
}

// Lv1 ...
func (s *Context) Lv1(message ...interface{}) {
	s.Logger.Info(s.toContextLogger("Lv1"), "", formatLogs(message...)...)
}

// Lv2 ...
func (s *Context) Lv2(message ...interface{}) time.Time {
	s.Logger.Info(s.toContextLogger("Lv2"), "", formatLogs(message...)...)
	return time.Now()
}

// Lv3 ...
func (s *Context) Lv3(startProcessTime time.Time, message ...interface{}) {
	stop := time.Now()

	msg := formatLogs(message...)
	msg = append(msg, Logger.ToField("_process_time", stop.Sub(startProcessTime).Nanoseconds()/1000000))

	s.Logger.Info(s.toContextLogger("Lv3"), "", msg...)
}

// Lv4 ...
func (s *Context) Lv4(message ...interface{}) {
	stop := time.Now()
	rt := stop.Sub(s.RequestTime).Nanoseconds() / 1000000

	msg := formatLogs(message...)
	msg = append(msg, Logger.ToField("_response_time", rt))
	msg = append(msg, Logger.ToField("_response_code", s.ResponseCode))

	s.Logger.Info(s.toContextLogger("Lv4"), "", msg...)
}

// formatLogs ...
func formatLogs(message ...interface{}) (logRecord []Logger.Field) {
	for index, msg := range message {
		logRecord = append(logRecord, Logger.ToField("_message_"+cast.ToString(index), msg))
	}

	return
}

// toContextLogger ...
func (s *Context) toContextLogger(tag string) (ctx context.Context) {
	ctxVal := Logger.Context{
		ServiceName:    s.AppName,
		ServiceVersion: s.AppVersion,
		ServicePort:    s.Port,
		XRequestID:     s.XRequestID,
		XAgent:         s.XAgent,
		Tag:            tag,
		ReqMethod:      s.Method,
		ReqURI:         s.URL,
		AdditionalData: s.Map.Items(),
		Error:          s.ErrorMessage,
	}
	if tag == "Lv4" {
		ctxVal.Request = Logger.ToField("req", s.Request)
		ctxVal.Response = Logger.ToField("resp", s.Response)
	}

	ctx = Logger.InjectCtx(context.Background(), ctxVal)
	return
}
