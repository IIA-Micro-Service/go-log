package log

import (
	"github.com/IIA-Micro-Service/go-log/adapter"
	loggerAdapter "github.com/IIA-Micro-Service/go-log/adapter/logrus"
	"github.com/IIA-Micro-Service/go-log/config"
	"github.com/IIA-Micro-Service/go-log/tracer"
)

func Init(config *config.Config) adapter.Logger {
	var loggerHandler adapter.Logger
	if "" == config.Type {
		config.Type = "logrus"
	}

	if "logrus" == config.Type {
		loggerHandler = loggerAdapter.NewLogrusWrapper()
	}

	// 再初始化Log组件的同时，去初始化Opentelemetry组件
	tracer.NewTracerWrapper()

	return loggerHandler
}
