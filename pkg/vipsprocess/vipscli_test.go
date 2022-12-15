package vipsprocess

import (
	"testing"
)

func Test_SaveWebp(t *testing.T) {
	SaveImage(&ProcessCtx{
		Quality: 50,
		Input:   "/Users/zhangyan/Documents/images/color.jpg",
	})
}
