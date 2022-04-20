package adapter

/*
 * @desc : 先定义好interface抽象接口，其他所有的三方开源组件都必须遵守并实现该接口
 */
type Fields map[string]interface{}
type Logger interface {
	WithField(key string, value interface{}) Logger
	WithFields(fields Fields) Logger
	Trace(args ...interface{})
	Tracef(format string, args ...interface{})
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
}
