package processcli

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/byyam/image-compare/pkg/logger"
)

// https://ffmpeg.org/ffmpeg-filters.html#libvmaf
func GetVMAF(opt *ProcessCtx) {
	cmd := fmt.Sprintf("ffmpeg -i %s -i %s -lavfi libvmaf='feature=name=ciede' -f null - 2>&1 | grep -oE \"VMAF score: \\d+.\\d+\" | grep -oE \"\\d+.\\d+\"", opt.Input, opt.Output)
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		logger.Warnf("exec[%s] failed", cmd)
		return
	}
	outline := strings.TrimSpace(string(out))
	if s, err := strconv.ParseFloat(outline, 64); err == nil {
		opt.VMAF = s
	} else {
		logger.Warnf("get vmaf float failed:[%s][%v]", outline, err)
	}
	logger.Debugf("exec[%s] done", cmd)
}

// https://ffmpeg.org/ffmpeg-filters.html#ssim
func GetSsim(opt *ProcessCtx) {
	cmd := fmt.Sprintf("ffmpeg -i %s -i %s -lavfi \"ssim\" -f null - 2>&1 | grep \"SSIM\" | grep -oE \"All:\\d+.\\d+\" | grep -oE \"\\d+.\\d+\"", opt.Input, opt.Output)
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		logger.Warnf("exec[%s] failed", cmd)
		return
	}
	outline := strings.TrimSpace(string(out))
	if s, err := strconv.ParseFloat(outline, 64); err == nil {
		opt.SSIM = s
	} else {
		logger.Warnf("get ssim float failed:[%s][%v]", outline, err)
	}
	logger.Debugf("exec[%s] done", cmd)
}

func GetPsnr(opt *ProcessCtx) {
	cmd := fmt.Sprintf("ffmpeg -i %s -i %s -lavfi \"psnr\" -f null - 2>&1 | grep \"PSNR\" | grep -oE \"average:\\d+.\\d+\" | grep -oE \"\\d+.\\d+\"", opt.Input, opt.Output)
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		logger.Warnf("exec[%s] failed", cmd)
		return
	}
	outline := strings.TrimSpace(string(out))
	if s, err := strconv.ParseFloat(outline, 64); err == nil {
		opt.PSNR = s
	} else {
		logger.Warnf("get psnr float failed:[%s][%v]", outline, err)
	}
	logger.Debugf("exec[%s] done", cmd)
}
