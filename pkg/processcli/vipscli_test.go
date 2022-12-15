package processcli

import (
	"testing"
)

func Test_SaveWebp(t *testing.T) {
	if err := SaveImage(&ProcessCtx{
		Quality: 50,
		Input:   "/Users/zhangyan/Documents/images/color.jpg",
	}); err != nil {
		t.Fatalf("save image failed:%v", err)
	}
}
