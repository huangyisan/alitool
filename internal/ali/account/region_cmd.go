package account

import (
	. "alitool/internal/pkg/mylog"
)

func ListRegion() {
	for k, v := range getRegionList() {
		LoggerNoT.Info(k, v)
	}
}
