package config

import (
	slogmulti "github.com/samber/slog-multi"
	"gopkg.in/natefinch/lumberjack.v2"
	"log/slog"
	"os"
)

var logger *slog.Logger

func Set(config *App) {
	logger = slog.New(
		slog.NewTextHandler(os.Stderr, nil),
	)

	if config.Env == "development" {
		logRotate := &lumberjack.Logger{
			Filename:   "log/app.log",
			MaxSize:    100,
			MaxBackups: 3,
			MaxAge:     28,
			Compress:   true,
		}
		logger = slog.New(
			slogmulti.Fanout(
				slog.NewJSONHandler(logRotate, nil),
				slog.NewTextHandler(os.Stderr, nil),
			),
		)
	}
	slog.SetDefault(logger)
}
