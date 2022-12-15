package vipsprocess

import (
	"fmt"
	"os/exec"
	"path"
	"strings"

	"github.com/byyam/image-compare/pkg/logger"
)

func getSuffix(opt *ProcessCtx) string {
	switch opt.TargetFormat {
	case FormatWebp:
		return ".webp"
	case FormatHeif:
		return ".heif"
	case FormatAvif:
		return ".avif"
	default:
		logger.Fatalf("unsupported format:%d", opt.TargetFormat)
	}
	return ""
}

func handleOutput(opt *ProcessCtx) {
	if opt.Output == "" {
		suffix := getSuffix(opt)
		fileSuffix := path.Ext(opt.Input)
		filenameOnly := strings.TrimSuffix(opt.Input, fileSuffix)
		opt.Output = filenameOnly + suffix
	}
}

func execCmd(cmd string) error {
	_, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		logger.Warnf("exec[%s] failed", cmd)
		return err
	}
	logger.Debugf("exec[%s] done", cmd)
	return nil
}

func handleByFormat(opt *ProcessCtx) string {
	handleOutput(opt)

	switch opt.TargetFormat {
	case FormatWebp:
		return fmt.Sprintf("vips webpsave -Q %d %s %s", opt.Quality, opt.Input, opt.Output)
	case FormatHeif:
		return fmt.Sprintf("vips heifsave -Q %d %s %s", opt.Quality, opt.Input, opt.Output)
	case FormatAvif:
		return fmt.Sprintf("vips heifsave -Q %d %s %s", opt.Quality, opt.Input, opt.Output)
	default:
		logger.Fatalf("unsupported format:%d", opt.TargetFormat)
	}
	return ""
}

func SaveImage(opt *ProcessCtx) error {
	cmd := handleByFormat(opt)
	if opt.Quality == 0 {
		logger.Warnf("Q is 0")
	}

	if err := execCmd(cmd); err != nil {
		return err
	}
	return nil
}
