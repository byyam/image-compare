package processcli

import (
	"os"
	"path"
	"path/filepath"
	"strings"

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
	// input context
	TargetFormat FormatType
	Quality      int
	Input        string
	Output       string
	// result context
	OutputFileInfo os.FileInfo
	SSIM           float64
	DSSIM          float64
	PSNR           float64
	VMAF           float64
}

func Process(ctx *ProcessCtx, options ...Option) error {
	opts := &Opts{
		EnableVMAF: false,
		EnablePSNR: false,
		EnableSSIM: false,
	}
	for _, opt := range options {
		opt(opts)
	}
	if err := SaveImage(ctx); err != nil {
		logger.Errorf("process failed:%v", err)
		return err
	}
	if err := GetFileInfo(ctx); err != nil {
		logger.Errorf("get output fileinfo failed:%v", err)
		return err
	}
	if opts.EnableVMAF {
		GetVMAF(ctx)
	}
	if opts.EnableSSIM {
		GetSsimIM(ctx)
	}
	if opts.EnableDSSIM {
		GetDSsimIM(ctx)
	}
	if opts.EnablePSNR {
		GetPsnrIM(ctx)
	}
	return nil
}

func GetFileInfo(ctx *ProcessCtx) error {
	fileInfo, err := os.Stat(ctx.Output)
	if err != nil {
		return err
	}
	ctx.OutputFileInfo = fileInfo
	return nil
}

type Opts struct {
	EnableSSIM  bool
	EnableDSSIM bool
	EnablePSNR  bool
	EnableVMAF  bool
}

type Option func(opts *Opts)

func WithSSIM(v bool) Option {
	return func(opts *Opts) {
		opts.EnableSSIM = v
	}
}

func WithDSSIM(v bool) Option {
	return func(opts *Opts) {
		opts.EnableDSSIM = v
	}
}

func WithPSNR(v bool) Option {
	return func(opts *Opts) {
		opts.EnablePSNR = v
	}
}

func WithVMAF(v bool) Option {
	return func(opts *Opts) {
		opts.EnableVMAF = v
	}
}

func GetFilePathAndSuffix(filepath string) (string, string) {
	fileSuffix := path.Ext(filepath)
	filenameOnly := strings.TrimSuffix(filepath, fileSuffix)
	return filenameOnly, fileSuffix
}

func WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix)
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}
		return nil
	})
	return files, err
}
