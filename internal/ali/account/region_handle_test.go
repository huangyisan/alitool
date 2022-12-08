package account

import (
	"github.com/agiledragon/gomonkey/v2"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_getRegionList(t *testing.T) {
	convey.Convey("Test_getRegionList", t, func() {
		convey.Convey("Patch regionNameMapping", func() {
			patches := gomonkey.ApplyGlobalVar(&regionNameMapping, map[region]string{"cn_shanghai_patched": "上海"})
			defer patches.Reset()
			regionNameMappingMock := getRegionList()
			convey.So(regionNameMappingMock, convey.ShouldResemble, regionNameMapping)
		})
	})
}
