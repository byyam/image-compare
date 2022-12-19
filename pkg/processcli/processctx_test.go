package processcli

import (
	"testing"
)

func TestProcess(t *testing.T) {
	ctx := &ProcessCtx{
		Quality:      50,
		TargetFormat: FormatWebp,
		Input:        "/Users/zhangyan/Documents/images/color.jpg",
	}
	if err := Process(ctx, WithSSIM(true), WithPSNR(true), WithVMAF(true)); err != nil {
		t.Fatalf("process failed:%v", err)
	}
	t.Logf("ctx:%+v,out size:%d", ctx, ctx.OutputFileInfo.Size())
}

func TestGetFilePathAndSuffix(t *testing.T) {
	filepath := "/Users/zhangyan/Documents/images/color.jpg"
	file, suffix := GetFilePathAndSuffix(filepath)
	t.Logf("file=%s, suffix=%s", file, suffix)
}
