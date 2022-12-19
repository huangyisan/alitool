package account

import (
	"alitool/internal/pkg/mylog"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestListRegion(t *testing.T) {
	wantKey01 := region("region01")
	wantKey02 := region("region02")
	wantValue01 := "region01_v"
	wantValue02 := "region02"
	convey.Convey("mock accountMap", t, func() {
		patches := gomonkey.ApplyGlobalVar(&regionNameMapping, map[region]string{
			wantKey01: wantValue01,
			wantKey02: wantValue02,
		})
		defer patches.Reset()
		convey.Convey("Capture log print", func() {
			res := mylog.CaptureLogInfo(ListRegion)
			convey.So(res, convey.ShouldContainSubstring, string(wantKey01))
			convey.So(res, convey.ShouldContainSubstring, string(wantKey02))
			convey.So(res, convey.ShouldContainSubstring, wantValue01)
			convey.So(res, convey.ShouldContainSubstring, wantValue02)
		})
	})
}
