package logrus

import (
	"context"
	"github.com/IIA-Micro-Service/go-log/adapter"
	"github.com/IIA-Micro-Service/go-log/adapter/logrus/hook"
	"github.com/IIA-Micro-Service/go-log/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
	"io/ioutil"
	"log"
	"time"
)

/*
 * @desc : 首先，请熟悉logrus用法与logrus基本概念
 *         logrus在直接New时候返回的是Log对象
 *         logrus在使用fileds后返回的是Entry对象
 *         但是说到底，最终都是通过Entry对象进行地日志写入
 */

/*
 * logrus的wrapper
 */
type logrusWrapper struct {
	handler *logrus.Logger // logrus对象句柄
	//tracerProvider trace.TracerProvider
	//tracer         trace.Tracer // tracer对象
}

func (lw *logrusWrapper) WithContext(ctx context.Context) adapter.Logger {
	//lw.handler.WithContext()
	return newLogrusEntryWrapper(lw.handler.WithContext(ctx))
}

/*
func (lw *logrusWrapper) GetTracerProvider() trace.TracerProvider {
	//return lw.handler.WithField(key, value)
	//return newLogrusEntryWrapper(lw.handler.WithField(key, value))
	return lw.tracerProvider
}
func (lw *logrusWrapper) GetTracer() trace.Tracer {
	//return lw.handler.WithField(key, value)
	//return newLogrusEntryWrapper(lw.handler.WithField(key, value))
	return lw.tracer
}
*/
func (lw *logrusWrapper) WithField(key string, value interface{}) adapter.Logger {
	//return lw.handler.WithField(key, value)
	return newLogrusEntryWrapper(lw.handler.WithField(key, value))
}
func (lw *logrusWrapper) WithFields(fields adapter.Fields) adapter.Logger {
	//return lw.handler.WithFields(fields)
	return newLogrusEntryWrapper(lw.handler.WithFields(logrus.Fields(fields)))
}
func (lw *logrusWrapper) Trace(args ...interface{}) {
	lw.handler.Trace(args)
}
func (lw *logrusWrapper) Tracef(format string, args ...interface{}) {
	lw.handler.Trace(format, args)
}
func (lw *logrusWrapper) Debug(args ...interface{}) {
	lw.handler.Debug(args)
}
func (lw *logrusWrapper) Debugf(format string, args ...interface{}) {
	lw.handler.Debugf(format, args)
}
func (lw *logrusWrapper) Info(args ...interface{}) {
	lw.handler.Info(args)
}
func (lw *logrusWrapper) Infof(format string, args ...interface{}) {
	lw.handler.Infof(format, args)
}
func (lw *logrusWrapper) Warn(args ...interface{}) {
	lw.handler.Warn(args)
}
func (lw *logrusWrapper) Warnf(format string, args ...interface{}) {
	lw.handler.Warnf(format, args)
}
func (lw *logrusWrapper) Error(args ...interface{}) {
	lw.handler.Error(args)
}
func (lw *logrusWrapper) Errorf(format string, args ...interface{}) {
	lw.handler.Errorf(format, args)
}
func (lw *logrusWrapper) Fatal(args ...interface{}) {
	lw.handler.Fatal(args)
}
func (lw *logrusWrapper) Fatalf(format string, args ...interface{}) {
	lw.handler.Fatalf(format, args)
}
func (lw *logrusWrapper) Panic(args ...interface{}) {
	lw.handler.Panic(args)
}
func (lw *logrusWrapper) Panicf(format string, args ...interface{}) {
	lw.handler.Panicf(format, args)
}

/*
 * logrus.Entry的wrapper
 */
type logrusEntryWrapper struct {
	entryHandler   *logrus.Entry
	tracerProvider trace.TracerProvider
	tracer         trace.Tracer // tracer对象
}

func (lew *logrusEntryWrapper) WithContext(ctx context.Context) adapter.Logger {
	//lew.entryHandler.WithContext(ctx)
	return newLogrusEntryWrapper(lew.entryHandler.WithContext(ctx))
}
func (lew *logrusEntryWrapper) GetTracerProvider() trace.TracerProvider {
	//return lw.handler.WithField(key, value)
	//return newLogrusEntryWrapper(lw.handler.WithField(key, value))
	return lew.tracerProvider
}
func (lew *logrusEntryWrapper) GetTracer() trace.Tracer {
	//return lw.handler.WithField(key, value)
	//return newLogrusEntryWrapper(lw.handler.WithField(key, value))
	return lew.tracer
}
func (lew *logrusEntryWrapper) WithField(key string, value interface{}) adapter.Logger {
	//return lew.entryHandler.WithField(key, value)
	return newLogrusEntryWrapper(lew.entryHandler.WithField(key, value))
}
func (lew *logrusEntryWrapper) WithFields(fields adapter.Fields) adapter.Logger {
	//return lw.handler.WithFields(fields)
	return newLogrusEntryWrapper(lew.entryHandler.WithFields(logrus.Fields(fields)))
}
func (lew *logrusEntryWrapper) Trace(args ...interface{}) {
	lew.entryHandler.Trace(args)
}
func (lew *logrusEntryWrapper) Tracef(format string, args ...interface{}) {
	lew.entryHandler.Trace(format, args)
}
func (lew *logrusEntryWrapper) Debug(args ...interface{}) {
	lew.entryHandler.Debug(args)
}
func (lew *logrusEntryWrapper) Debugf(format string, args ...interface{}) {
	lew.entryHandler.Debugf(format, args)
}
func (lew *logrusEntryWrapper) Info(args ...interface{}) {
	lew.entryHandler.Info(args)
}
func (lew *logrusEntryWrapper) Infof(format string, args ...interface{}) {
	lew.entryHandler.Infof(format, args)
}
func (lew *logrusEntryWrapper) Warn(args ...interface{}) {
	lew.entryHandler.Warn(args)
}
func (lew *logrusEntryWrapper) Warnf(format string, args ...interface{}) {
	lew.entryHandler.Warnf(format, args)
}
func (lew *logrusEntryWrapper) Error(args ...interface{}) {
	lew.entryHandler.Error(args)
}
func (lew *logrusEntryWrapper) Errorf(format string, args ...interface{}) {
	lew.entryHandler.Errorf(format, args)
}
func (lew *logrusEntryWrapper) Fatal(args ...interface{}) {
	lew.entryHandler.Fatal(args)
}
func (lew *logrusEntryWrapper) Fatalf(format string, args ...interface{}) {
	lew.entryHandler.Fatalf(format, args)
}
func (lew *logrusEntryWrapper) Panic(args ...interface{}) {
	lew.entryHandler.Panic(args)
}
func (lew *logrusEntryWrapper) Panicf(format string, args ...interface{}) {
	lew.entryHandler.Panicf(format, args)
}

var _ adapter.Logger = &logrusEntryWrapper{}
var _ adapter.Logger = &logrusWrapper{}

func newLogrusEntryWrapper(entry *logrus.Entry) *logrusEntryWrapper {
	return &logrusEntryWrapper{
		entryHandler: entry,
	}
}

func NewLogrusWrapper(config *config.Config) *logrusWrapper {

	// 日志文件
	logFile := config.LogDir + config.LogFileName
	//logFileHandler, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//if err != nil {}

	// 配置日志轮换
	rotateWriter, err := rotatelogs.New(
		logFile+".%Y%m%d",
		rotatelogs.WithLinkName(logFile), // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(time.Duration(config.MaxLife)*time.Second),            // 文件最大保存时间
		rotatelogs.WithRotationTime(time.Duration(config.RotationTime)*time.Second), // 日志切割时间间隔
	)
	if err != nil {
		log.Fatalf("rotatelogs.New err: %+v", err)
	}

	// 通过hook方式，将rotateWriter hook到logrus上去
	rotateWriterHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.TraceLevel: rotateWriter,
		logrus.DebugLevel: rotateWriter,
		logrus.InfoLevel:  rotateWriter,
		logrus.WarnLevel:  rotateWriter,
		logrus.ErrorLevel: rotateWriter,
		logrus.FatalLevel: rotateWriter,
		logrus.PanicLevel: rotateWriter,
	}, &logrus.JSONFormatter{})

	logrusHandler := logrus.New()
	//logger.Out = ioutil.Discard
	// 关闭logrus日志在屏幕stdout上的输出
	logrusHandler.SetOutput(ioutil.Discard)
	//logger.SetOutput(writer)
	//logger.Hooks.Add(lfsHook)
	logrusHandler.SetLevel(logrus.Level(config.LogLevel))
	logrusHandler.AddHook(hook.NewTracerHook())
	logrusHandler.AddHook(rotateWriterHook)

	/*
		logrusHandler.WithFields(logrus.Fields{
			"traceid" : time.Now(),
		})
	*/

	return &logrusWrapper{
		handler: logrusHandler,
	}
}
