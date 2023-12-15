package LogTest

import (
	"context"
	"gopkg.in/natefinch/lumberjack.v2"
	"log/slog"
	"net"
	"os"
	"testing"
)

func TestSLog(t *testing.T) {
	slog.Info("my first slog msg", "greeting", "hello, slog")
}

func TestSLogTextHandler(t *testing.T) {

	// time=2023-12-12T16:13:48.810+08:00 level=INFO msg="hello world" name="tony smile"
	// time=2023-12-12T16:13:48.810+08:00 level=ERROR msg=oops err="use of closed network connection" status=500
	h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{AddSource: true})
	l := slog.New(h)
	l.Info("hello world", "name", "tony smile")
	l.Error("oops", "err", net.ErrClosed, "status", 500)
}

func TestSLogJsonHandler(t *testing.T) {
	// {"time":"2023-12-12T16:13:48.810726+08:00","level":"INFO","msg":"hello world","name":"tony"}
	// {"time":"2023-12-12T16:13:48.810758+08:00","level":"ERROR","msg":"oops","err":"use of closed network connection","status":500}
	h := slog.NewJSONHandler(os.Stderr, nil)
	l := slog.New(h)
	l.Info("hello world", "name", "tony smile")
	l.Error("oops", "err", net.ErrClosed, "status", 500)
}

func TestSLogFile(t *testing.T) {
	logFile := &lumberjack.Logger{
		Filename:   "./logs/app.log",
		MaxSize:    5, // MB
		MaxBackups: 3,
		MaxAge:     7, // days
		Compress:   true,
		LocalTime:  true,
	}

	for {
		logger := slog.New(slog.NewTextHandler(logFile, nil))
		logger.Info("hello world", "name", "tony")
		logger.Error("oops", "err", net.ErrClosed, "status", 500)

		slog.SetDefault(logger)
		slog.Info("Info message")

		logger.LogAttrs(
			context.Background(),
			slog.LevelInfo,
			"incoming request",
			slog.String("method", "GET"),
			slog.Int("time_taken_ms", 158),
			slog.String("path", "/hello/world?q=search"),
			slog.Int("status", 200),
			slog.String(
				"user_agent",
				"Googlebot/2.1 (+http://www.google.com/bot.html)",
			),
		)

		logger.LogAttrs(
			context.Background(),
			slog.LevelInfo,
			"image uploaded",
			slog.Int("id", 23123),
			slog.Group("properties",
				slog.Int("width", 4000),
				slog.Int("height", 3000),
				slog.String("format", "jpeg"),
			),
		)

	}
}
