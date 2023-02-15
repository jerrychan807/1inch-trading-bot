package lib

import (
	"github.com/sirupsen/logrus"
	"os"
)

var (
	Logger = logrus.New()
)

func init() {
	// 为当前logrus实例设置消息的输出,同样地,
	// 可以设置logrus实例的输出到任意io.writer
	Logger.Out = os.Stdout
	Logger.SetLevel(logrus.DebugLevel)
}
