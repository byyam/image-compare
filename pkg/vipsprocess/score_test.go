package vipsprocess

import (
	"log"
	"os"
	"testing"

	"github.com/byyam/image-compare/pkg/logger"
)

func TestMain(m *testing.M) {
	logger.SetLogLevel(logger.LOGLEVEL_DEBUG)

	log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")

	os.Exit(exitVal)
}

func Test_GetVMAF(t *testing.T) {
	ctx := &ProcessCtx{
		Input:  "/Users/zhangyan/Documents/images/color.jpg",
		Output: "/Users/zhangyan/Documents/images/color.webp",
	}
	GetVMAF(ctx)
	t.Logf("ctx:%+v", ctx)
}

func Test_GetSsim(t *testing.T) {
	ctx := &ProcessCtx{
		Input:  "/Users/zhangyan/Documents/images/color.jpg",
		Output: "/Users/zhangyan/Documents/images/color.webp",
	}
	GetSsim(ctx)
	t.Logf("ctx:%+v", ctx)
}

func Test_GetPsnr(t *testing.T) {
	ctx := &ProcessCtx{
		Input:  "/Users/zhangyan/Documents/images/color.jpg",
		Output: "/Users/zhangyan/Documents/images/color.webp",
	}
	GetPsnr(ctx)
	t.Logf("ctx:%+v", ctx)
}
