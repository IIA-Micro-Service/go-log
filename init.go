package log

import (
	"github.com/IIA-Micro-Service/go-log/adapter"
	loggerAdapter "github.com/IIA-Micro-Service/go-log/adapter/logrus"
	"github.com/IIA-Micro-Service/go-log/config"
)

func Init(config *config.Config) adapter.Logger {
	var loggerHandler adapter.Logger
	if "" == config.Type {
		config.Type = "logrus"
	}

	if "logrus" == config.Type {
		loggerHandler = loggerAdapter.NewLogrusWrapper()
	}
	return loggerHandler
}
