package logs

import (
	"os"

	"github.com/google/logger"
)

/* 日志记录对象 */

var (
	lf *os.File
)

// New 初始化日志记录对象
func New(logPath, name string, verbose, systemLog bool) (err error) {
	lf, err = os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		logger.Fatalf("打开日志文件失败: %v", err)
		return err
	}
	// 初始化日志对象
	logger.Init(name, verbose, systemLog, lf)
	return nil
}

// Close 关闭日志对象
func Close() error {
	return lf.Close()
}
