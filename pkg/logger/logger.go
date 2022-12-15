package logger

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
)

type LOGLEVEL int

const (
	LOGLEVEL_UNKNOWN LOGLEVEL = iota
	LOGLEVEL_DEBUG
	LOGLEVEL_INFO
	LOGLEVEL_WARN
	LOGLEVEL_ERROR
	LOGLEVEL_FATAL
)

var (
	errorColor = color.New(color.FgRed)
	warnColor  = color.New(color.FgYellow)
	debugColor = color.New(color.FgGreen)
	infoColor  = color.New(color.FgBlue)
)

var (
	traceId  string
	logger   *log.Logger
	logLevel LOGLEVEL
)

func SetLogLevel(level LOGLEVEL) {
	logLevel = level
}

func init() {
	// 命令行工具的log封装，对log库简单封装，打印traceId来区别每次调用
	traceId = fmt.Sprintf("%d-%d", os.Getpid(), time.Now().Nanosecond())
	logger = log.New(os.Stdout, "["+traceId+"]", log.LstdFlags)
	logLevel = LOGLEVEL_INFO
}

func Debugf(format string, args ...interface{}) {
	if logLevel > LOGLEVEL_DEBUG {
		return
	}
	logger.Printf(debugColor.Sprintf("[DBG] "+format, args...))
}

func Infof(format string, args ...interface{}) {
	if logLevel > LOGLEVEL_INFO {
		return
	}
	logger.Printf(infoColor.Sprintf("[INF]"+format, args...))
}

func Warnf(format string, args ...interface{}) {
	if logLevel > LOGLEVEL_WARN {
		return
	}
	logger.Printf(warnColor.Sprintf("[WRN] "+format, args...))
}

func Errorf(format string, args ...interface{}) {
	if logLevel > LOGLEVEL_ERROR {
		return
	}
	logger.Printf(errorColor.Sprintf("[ERR] "+format, args...))
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(errorColor.Sprintf("[FTL] "+format, args...))
}
