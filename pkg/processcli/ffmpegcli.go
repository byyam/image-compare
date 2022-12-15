package processcli

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/byyam/image-compare/pkg/logger"
)

// https://ffmpeg.org/ffmpeg-filters.html#libvmaf
func GetVMAF(opt *ProcessCtx) {
	cmd := fmt.Sprintf("ffmpeg -i %s -i %s -lavfi libvmaf='feature=name=ciede' -f null - 2>&1 | grep \"VMAF score\"", opt.Input, opt.Output)
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		logger.Warnf("exec[%s] failed", cmd)
		return
	}
	outline := strings.TrimSpace(string(out))
	results := strings.SplitAfter(outline, "VMAF score: ")
	if len(results) == 0 {
		logger.Warnf("get vmaf failed:[%s]", outline)
		return
	}
	if s, err := strconv.ParseFloat(results[len(results)-1], 64); err == nil {
		opt.VMAF = s
	} else {
		logger.Warnf("get vmaf float failed:[%s][%v]", outline, err)
	}
	logger.Debugf("exec[%s] done", cmd)
}

// https://ffmpeg.org/ffmpeg-filters.html#ssim
func GetSsim(opt *ProcessCtx) {
	cmd := fmt.Sprintf("ffmpeg -i %s -i %s -lavfi \"ssim\" -f null - 2>&1 | grep \"SSIM\"", opt.Input, opt.Output)
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		logger.Warnf("exec[%s] failed", cmd)
		return
	}
	outline := strings.TrimSpace(string(out))
	reg := regexp.MustCompile(`All:\d+\.\d+`)
	if reg == nil {
		logger.Warnf("reg ssim failed:[%s]", outline)
		return
	}
	results := reg.FindAllString(outline, -1)
	if len(results) == 0 {
		logger.Warnf("get ssim failed:[%s]%+v", outline, results)
		return
	}
	logger.Debugf("results:%+v", results[0])
	scores := strings.Split(results[0], ":")
	if len(scores) != 2 {
		logger.Warnf("get ssim failed:[%s]", outline)
		return
	}
	if s, err := strconv.ParseFloat(scores[1], 64); err == nil {
		opt.SSIM = s
	} else {
		logger.Warnf("get ssim float failed:[%s][%v]", outline, err)
	}
	logger.Debugf("exec[%s] done", cmd)
}

func GetPsnr(opt *ProcessCtx) {
	cmd := fmt.Sprintf("ffmpeg -i %s -i %s -lavfi \"psnr\" -f null - 2>&1 | grep \"PSNR\"", opt.Input, opt.Output)
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		logger.Warnf("exec[%s] failed", cmd)
		return
	}
	outline := strings.TrimSpace(string(out))
	reg := regexp.MustCompile(`average:\d+\.\d+`)
	if reg == nil {
		logger.Warnf("reg psnr failed:[%s]", outline)
		return
	}
	results := reg.FindAllString(outline, -1)
	if len(results) == 0 {
		logger.Warnf("get psnr failed:[%s]%+v", outline, results)
		return
	}
	logger.Debugf("results:%+v", results[0])
	scores := strings.Split(results[0], ":")
	if len(scores) != 2 {
		logger.Warnf("get psnr failed:[%s]", outline)
		return
	}
	if s, err := strconv.ParseFloat(scores[1], 64); err == nil {
		opt.PSNR = s
	} else {
		logger.Warnf("get psnr float failed:[%s][%v]", outline, err)
	}
	logger.Debugf("exec[%s] done", cmd)
}
