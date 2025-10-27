package logger

import (
	"go.uber.org/zap"
)

var logger *zap.Logger

func InitLogger(level string, development bool) error {
	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return err
	}

	var cfg zap.Config
	if development {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
	}
	cfg.Level = lvl
	zl, err := cfg.Build()
	if err != nil {
		return err
	}
	logger = zl
	return nil
}

func Logger() *zap.Logger {
	return logger
}
