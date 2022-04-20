package hook

import (
	"github.com/sirupsen/logrus"
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

	return nil
}
