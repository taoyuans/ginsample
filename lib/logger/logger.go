package logger

import (
	"context"
	"io"
	"net"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"gopkg.in/natefinch/lumberjack.v2"
)

var once sync.Once

var log zerolog.Logger

type LogContext struct {
	log *zerolog.Logger

	TraceId    string      `json:"trace_id,omitempty"`
	Method     string      `json:"method,omitempty"`
	RequestUri string      `json:"request_uri,omitempty"`
	UserAgent  string      `json:"user_agent,omitempty"`
	Host       string      `json:"host,omitempty"`
	Ip         string      `json:"ip,omitempty"`
	Referer    string      `json:"referer,omitempty"`
	Path       string      `json:"path,omitempty"`
	Params     interface{} `json:"params,omitempty"`
	Status     int         `json:"status,omitempty"`
}

func Get(filename string, maxSize, maxBackups, maxAge int, logLevel int8) *zerolog.Logger {
	once.Do(func() {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		zerolog.TimeFieldFormat = time.RFC3339Nano

		var output io.Writer = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}
		fileLogger := &lumberjack.Logger{
			Filename:   filename,
			MaxSize:    maxSize,
			MaxBackups: maxBackups,
			MaxAge:     maxAge,
			Compress:   true,
		}

		output = zerolog.MultiLevelWriter(os.Stderr, fileLogger)

		var gitRevision string

		buildInfo, ok := debug.ReadBuildInfo()
		if ok {
			for _, v := range buildInfo.Settings {
				if v.Key == "vcs.revision" {
					gitRevision = v.Value
					break
				}
			}
		}

		log = zerolog.New(output).
			Level(zerolog.Level(logLevel)).
			With().
			Timestamp().
			Str("git_revision", gitRevision).
			Str("go_version", buildInfo.GoVersion).
			Logger()
	})

	return &log
}

func GetLogger(ctx context.Context) *LogContext {
	v := ctx.Value("logContext")
	if v == nil {
		panic("Logger is not exist")
	}
	if logger, ok := v.(*LogContext); ok {
		return logger
	}
	panic("Logger is not exist")
}

func (logCtx *LogContext) Debug() *zerolog.Event {
	return logCtx.log.Debug().
		Str("trace_id", logCtx.TraceId).
		Str("method", logCtx.Method).
		Str("request_uri", logCtx.RequestUri).
		Str("user_agent", logCtx.UserAgent).
		Str("host", logCtx.Host).
		Str("ip", logCtx.Ip).
		Str("referer", logCtx.Referer).
		Str("path", logCtx.Path).
		Interface("params", logCtx.Params)
}

func (logCtx *LogContext) Info() *zerolog.Event {
	return logCtx.log.Info().
		Str("trace_id", logCtx.TraceId).
		Str("method", logCtx.Method).
		Str("request_uri", logCtx.RequestUri).
		Str("user_agent", logCtx.UserAgent).
		Str("host", logCtx.Host).
		Str("ip", logCtx.Ip).
		Str("referer", logCtx.Referer).
		Str("path", logCtx.Path).
		Interface("params", logCtx.Params)
}

func (logCtx *LogContext) Warn() *zerolog.Event {
	return logCtx.log.Warn().
		Str("trace_id", logCtx.TraceId).
		Str("method", logCtx.Method).
		Str("request_uri", logCtx.RequestUri).
		Str("user_agent", logCtx.UserAgent).
		Str("host", logCtx.Host).
		Str("ip", logCtx.Ip).
		Str("referer", logCtx.Referer).
		Str("path", logCtx.Path).
		Interface("params", logCtx.Params)
}

func (logCtx *LogContext) Error() *zerolog.Event {
	return logCtx.log.Error().
		Str("trace_id", logCtx.TraceId).
		Str("method", logCtx.Method).
		Str("request_uri", logCtx.RequestUri).
		Str("user_agent", logCtx.UserAgent).
		Str("host", logCtx.Host).
		Str("ip", logCtx.Ip).
		Str("referer", logCtx.Referer).
		Str("path", logCtx.Path).
		Interface("params", logCtx.Params)
}

func New(req *http.Request, fileName string, maxSize, maxBackups, maxAge int, logLevel int8) *LogContext {
	realIP := req.RemoteAddr
	if ip := req.Header.Get("X-Forwarded-For"); ip != "" {
		realIP = strings.Split(ip, ", ")[0]
	} else if ip := req.Header.Get("X-Real-IP"); ip != "" {
		realIP = ip
	} else {
		realIP, _, _ = net.SplitHostPort(realIP)
	}

	path := req.URL.Path
	if path == "" {
		path = "/"
	}

	liblogger := Get(fileName, maxSize, maxBackups, maxAge, logLevel)

	var params interface{}
	if req.Method == http.MethodGet {
		params = req.URL.Query()
	} else {
		// data, _ := ioutil.ReadAll(req.Body)
		// params = string(data)
	}

	logContext := &LogContext{
		log:        liblogger,
		TraceId:    req.Header.Get("TraceId"),
		Method:     req.Method,
		RequestUri: req.URL.RequestURI(),
		UserAgent:  req.UserAgent(),
		Host:       req.Host,
		Ip:         realIP,
		Referer:    req.Referer(),
		Path:       path,
		Params:     params,
	}
	return logContext
}
