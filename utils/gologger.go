package utils

import (
	"github.com/projectdiscovery/gologger/formatter"
	"github.com/projectdiscovery/gologger/levels"
	"go.uber.org/zap"
)

var _ formatter.Formatter = &GoLoggerFormatter{}

type GoLoggerFormatter struct {
	logger *zap.Logger
}

func ToGoLoggerFormatter(logger *zap.Logger) *GoLoggerFormatter {
	return &GoLoggerFormatter{logger: logger}
}

func map2fields(m map[string]string) []zap.Field {
	fields := make([]zap.Field, 0, len(m))
	for k, v := range m {
		if k != "label" {
			fields = append(fields, zap.String(k, v))
		}
	}
	return fields
}

func (f *GoLoggerFormatter) Format(event *formatter.LogEvent) ([]byte, error) {
	switch event.Level {
	case levels.LevelVerbose:
		f.logger.Debug(event.Message, map2fields(event.Metadata)...)
	case levels.LevelDebug:
		f.logger.Debug(event.Message, map2fields(event.Metadata)...)
	case levels.LevelInfo:
		f.logger.Info(event.Message, map2fields(event.Metadata)...)
	case levels.LevelWarning:
		f.logger.Warn(event.Message, map2fields(event.Metadata)...)
	case levels.LevelError:
		f.logger.Error(event.Message, map2fields(event.Metadata)...)
	case levels.LevelFatal:
		f.logger.Fatal(event.Message, map2fields(event.Metadata)...)
	}
	return nil, nil
}
