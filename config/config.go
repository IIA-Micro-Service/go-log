package config

import (
	"time"
)

type Config struct {
	Type         string        // 选用的三方日志库
	LogDir       string        // 日志目录名称
	LogFileName  string        // 日志文件名称
	MaxLife      time.Duration // 日志最长存活时间
	RotationTime time.Duration // 日志分割周期
	LogLevel     uint32        // 日志等级
	TracerConfig               // trace的配置
	MetricConfig               // metric的配置
}

type TracerConfig struct {
	OpenTracer bool // 是否打开Log组件自带的otel组件？
}

type MetricConfig struct {
}
