package hook

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
)

/*
 * @desc : logrus自定义hook插件
 */
type TracerHook struct{}

/**/
func (h *TracerHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// fire会被触发...
func (h *TracerHook) Fire(entry *logrus.Entry) error {
	//fmt.Printf("%+v", entry.Context)
	if nil == entry.Context {
		return nil
	}
	span := trace.SpanFromContext(entry.Context)
	fmt.Printf("traceid:%s, spanid:%s\n", span.SpanContext().TraceID(), span.SpanContext().SpanID())
	//entry.WithField("traceid", span.SpanContext().TraceID())
	//entry.WithField("spanid", span.SpanContext().SpanID())
	entry.Data["traceid"] = span.SpanContext().TraceID()
	entry.Data["spanid"] = span.SpanContext().SpanID()
	return nil
}

func NewTracerHook() *TracerHook {
	return &TracerHook{}
}
