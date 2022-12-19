package processcli

import (
	"testing"
)

func Test_SaveWebp(t *testing.T) {
	if err := SaveImage(&ProcessCtx{
		TargetFormat: FormatWebp,
		Quality:      50,
		Input:        "/Users/zhangyan/Documents/images/color.jpg",
	}); err != nil {
		t.Fatalf("save image failed:%v", err)
	}
}

func BenchmarkSaveImage100webp(b *testing.B) {
	benchmarkGenerate(100, FormatWebp, b)
}

func benchmarkGenerate(i int, format FormatType, b *testing.B) {
	for n := 0; n < b.N; n++ {
		generate(i, format)
	}
}

func generate(n int, format FormatType) {
	ctx := &ProcessCtx{
		TargetFormat: format,
		Quality:      50,
		Input:        "/Users/zhangyan/Documents/images/color.jpg",
	}
	for i := 0; i < n; i++ {
		_ = SaveImage(ctx)
	}
}
