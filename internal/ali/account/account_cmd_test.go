package account

import (
	"alitool/internal/pkg/mylog"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestListAccount(t *testing.T) {
	want01 := "account01"
	want02 := "account02"
	convey.Convey("mock accountMap", t, func() {
		patches := gomonkey.ApplyGlobalVar(&accountMap, map[string]map[string]string{
			want01: {},
			want02: {},
		})
		defer patches.Reset()
		convey.Convey("Capture log print", func() {
			res := mylog.CaptureLogInfo(ListAccount)
			convey.So(res, convey.ShouldContainSubstring, want01)
			convey.So(res, convey.ShouldContainSubstring, want02)
		})
	})
}
