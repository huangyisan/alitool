package mylog

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	LoggerNoT = NewLogWithNoTime()
)

type NoTimeLogOutPut struct {
}

func (f *NoTimeLogOutPut) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("%s", entry.Message)), nil
}

func NewLogWithNoTime() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&NoTimeLogOutPut{})
	return logger
}

func CaptureLogInfo(f func()) string {
	var buff bytes.Buffer
	LoggerNoT.SetOutput(&buff)
	defer func() { LoggerNoT.SetOutput(os.Stderr) }()
	f()
	return buff.String()
}
