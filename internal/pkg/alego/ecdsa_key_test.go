package alego

import (
	"testing"
)

func Test_make(t *testing.T) {
	p := makeECDSAKey()
	t.Logf("%#v", p)
}
