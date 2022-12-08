package account

import (
	"bytes"
	"fmt"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/smartystreets/goconvey/convey"
	"io"
	"log"
	"os"
	"testing"
)

func TestListAccount(t *testing.T) {
	patches := gomonkey.ApplyGlobalVar(&accountMap,
		map[string]map[string]string{"account01": {"ak": "ak01", "sk": "sk01"}})
	defer patches.Reset()
	want := "account01"
	buf := new(bytes.Buffer)
	log.SetOutput(buf)
	defer log.SetOutput(os.Stderr)
	ListAccount()
	convey.Convey("TestListAccount", t, func() {

		fmt.Println(buf.String())
		convey.So(buf.String(), convey.ShouldContainSubstring, want)

	})

}

func PrintLoop() {
	fmt.Println("hello world")
	return
}
func TestPrintLoop(t *testing.T) {
	// 定义需要查找的子字符串
	var want = "hello world"

	// 调用 PrintLoop 函数，并捕获输出
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	PrintLoop()
	fmt.Println(buf.String())
	t.Log(buf.String())

	// 使用 So 方法判断调用 PrintLoop 函数的执行结果中是否包含子字符串
	convey.Convey("TestPrintLoop", t, func() {
		convey.So(buf.String(), convey.ShouldContainSubstring, want)
	})
}

func readByte( /*...*/ ) {
	// ...
	err := io.EOF // force an error
	if err != nil {
		fmt.Println("ERROR")
		log.Print("Couldn't read first byte")
		return
	}
	// ...
}

func TestReadByte(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stderr)
	readByte()
	t.Log(buf.String())
}
