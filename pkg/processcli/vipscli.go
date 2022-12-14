package processcli

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/byyam/image-compare/pkg/logger"
)

func getSuffix(opt *ProcessCtx) (string, error) {
	switch opt.TargetFormat {
	case FormatWebp, FormatHeif, FormatAvif:
		return "." + opt.TargetFormat.String(), nil

	default:
		logger.Errorf("unsupported format:%d", opt.TargetFormat)
	}
	return "", errors.New("unsupported format")
}

func handleOutput(opt *ProcessCtx) error {
	if opt.Output == "" {
		suffix, err := getSuffix(opt)
		if err != nil {
			return err
		}
		filenameOnly, _ := GetFilePathAndSuffix(opt.Input)
		opt.Output = filenameOnly + suffix
	}
	return nil
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

func handleByFormat(opt *ProcessCtx) (string, error) {
	if err := handleOutput(opt); err != nil {
		return "", err
	}

	switch opt.TargetFormat {
	case FormatWebp:
		return fmt.Sprintf("vips webpsave -Q %d %s %s", opt.Quality, opt.Input, opt.Output), nil
	case FormatHeif:
		return fmt.Sprintf("vips heifsave -Q %d %s %s", opt.Quality, opt.Input, opt.Output), nil
	case FormatAvif:
		if opt.Effort == 0 {
			return fmt.Sprintf("vips heifsave -Q %d %s %s", opt.Quality, opt.Input, opt.Output), nil
		} else {
			return fmt.Sprintf("vips heifsave --effort %d -Q %d %s %s", opt.Effort, opt.Quality, opt.Input, opt.Output), nil
		}
	default:
		logger.Errorf("unsupported format:%d", opt.TargetFormat)
	}
	return "", errors.New("unsupported format")
}

func SaveImage(opt *ProcessCtx) error {
	cmd, err := handleByFormat(opt)
	if err != nil {
		return err
	}
	if opt.Quality == 0 {
		logger.Warnf("Q is 0")
	}

	if err := execCmd(cmd); err != nil {
		return err
	}
	return nil
}
