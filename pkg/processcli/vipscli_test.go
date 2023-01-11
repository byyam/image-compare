package processcli

import (
	"testing"
)

const (
	convertNum     = 10
	convertQuality = 50
)

func Test_SaveWebp(t *testing.T) {
	for i := 0; i < convertNum; i++ {
		if err := SaveImage(&ProcessCtx{
			TargetFormat: FormatWebp,
			Quality:      convertQuality,
			Input:        "/go/src/bimg/testdata/test.jpg",
		}); err != nil {
			t.Fatalf("save image failed:%v", err)
		}
	}
}

func Test_SaveAvif(t *testing.T) {
	for i := 0; i < convertNum; i++ {
		if err := SaveImage(&ProcessCtx{
			TargetFormat: FormatAvif,
			Quality:      convertQuality,
			Effort:       4,
			Input:        "/go/src/bimg/testdata/test.jpg",
		}); err != nil {
			t.Fatalf("save image failed:%v", err)
		}
	}
}

func Test_SaveHeif(t *testing.T) {
	for i := 0; i < convertNum; i++ {
		if err := SaveImage(&ProcessCtx{
			TargetFormat: FormatHeif,
			Quality:      convertQuality,
			Input:        "/go/src/bimg/testdata/test.jpg",
		}); err != nil {
			t.Fatalf("save image failed:%v", err)
		}
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
		Quality:      convertQuality,
		Input:        "/go/src/bimg/testdata/test.jpg",
	}
	for i := 0; i < n; i++ {
		_ = SaveImage(ctx)
	}
}
