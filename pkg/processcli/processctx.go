package processcli

import (
	"github.com/byyam/image-compare/pkg/logger"
)

type FormatType int

const (
	FormatWebp FormatType = iota + 1
	FormatHeif
	FormatAvif
)

func (p FormatType) String() string {
	switch p {
	case FormatWebp:
		return "webp"
	case FormatHeif:
		return "heif"
	case FormatAvif:
		return "avif"
	}
	return ""
}

type ProcessCtx struct {
	TargetFormat FormatType
	Quality      int
	Input        string
	Output       string
	FileSize     int
	SSIM         float64
	PSNR         float64
	VMAF         float64
}

func Process(opt *ProcessCtx) {
	if err := SaveImage(opt); err != nil {
		logger.Fatalf("process failed:%v", err)
		return
	}

	GetVMAF(opt)
	GetSsim(opt)
	GetPsnr(opt)
}
