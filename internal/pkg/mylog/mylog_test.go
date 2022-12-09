package mylog

import (
	"testing"
)

func Test_a(t *testing.T) {
	a := NewLogWithNoTime()
	a.Info("ccccc")
}

//func captureOutput(f func()) string {
//	var buf bytes.Buffer
//	log.SetOutput(&buf)
//	f()
//	log.SetOutput(os.Stderr)
//	return buf.String()
//}

func TestStdWrap(t *testing.T) {
	a := CaptureLogInfo(func() {
		LoggerNoT.Info("abc")
	})
	//assert.Equal(t, "abc", a)
}
