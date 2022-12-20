package processcli

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/byyam/image-compare/pkg/logger"
)

const (
	metricPsnr  = "PSNR"
	metricSsim  = "SSIM"
	metricDSsim = "DSSIM"
)

func GetPsnrIM(opt *ProcessCtx) {
	opt.PSNR = GetMetric(opt, metricPsnr)
}

func GetSsimIM(opt *ProcessCtx) {
	opt.SSIM = GetMetric(opt, metricSsim)
}

func GetDSsimIM(opt *ProcessCtx) {
	opt.DSSIM = GetMetric(opt, metricDSsim)
}

func GetMetric(opt *ProcessCtx, metric string) (value float64) {
	cmd := fmt.Sprintf("magick compare -metric %s %s %s /dev/null  2>&1 | awk -F ' ' '{print $1}'", metric, opt.Input, opt.Output)
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		logger.Warnf("exec[%s] failed:%v", cmd, err)
		return
	}
	outline := strings.TrimSpace(string(out))
	if s, err := strconv.ParseFloat(outline, 64); err == nil {
		value = s
	} else {
		logger.Warnf("get psnr float failed:[%s][%v]", outline, err)
	}
	logger.Debugf("exec[%s] done", cmd)
	return
}
