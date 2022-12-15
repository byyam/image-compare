package vipsprocess

import (
	"testing"
)

func TestProcess(t *testing.T) {
	ctx := &ProcessCtx{
		Quality:      50,
		TargetFormat: FormatWebp,
		Input:        "/Users/zhangyan/Documents/images/color.jpg",
	}
	Process(ctx)
	t.Logf("ctx:%+v", ctx)
}
