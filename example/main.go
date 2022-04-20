package main

import (
	"github.com/IIA-Micro-Service/go-log"
	"github.com/IIA-Micro-Service/go-log/adapter"
	"github.com/IIA-Micro-Service/go-log/config"
	"time"
)

func main() {

	logConfig := &config.Config{
		Type:         "logrus",
		LogDir:       "./",
		LogFileName:  "logrus.log",
		MaxLife:      time.Duration(7*86400) * time.Second,
		RotationTime: time.Duration(1*86400) * time.Second,
	}
	loggerHandler := log.Init(logConfig)
	loggerHandler.WithFields(adapter.Fields{
		"trace": time.Now(),
	}).Info("xoxo", "xoxoxoxoxoxoxo")

	//tp := otel.GetTracerProvider()
	//fmt.Printf("3:%+v\n", tp)

}
