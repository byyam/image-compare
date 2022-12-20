package processcli

import (
	"testing"
)

func TestGetMetric(t *testing.T) {
	ctx := &ProcessCtx{
		Input:  "/Users/zhangyan/Documents/images/color.jpg",
		Output: "/Users/zhangyan/Documents/images/color.webp",
	}
	GetPsnrIM(ctx)
	GetSsimIM(ctx)
	GetDSsimIM(ctx)
	t.Logf("ctx:%+v", ctx)
}
