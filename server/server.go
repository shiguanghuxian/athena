package server

import (
	"fmt"
	"log"
	"os"

	"github.com/google/logger"
	"github.com/shiguanghuxian/athena/modules/common"
	"github.com/shiguanghuxian/athena/modules/config"
	"github.com/shiguanghuxian/athena/modules/logs"
)

/* server服务，用于接收采集端上传的数据，转存到mq，同时做数据验证和权限验证 */

// ServerService 数据接收服务
type ServerService struct {
	config *config.ServerConfig
}

// New 创建一个数据接收服务
func New() *ServerService {
	serverService := new(ServerService)
	// 跟目录
	rootDir := common.GetRootDir()
	// 目录分隔符
	pathSeparator := string(os.PathSeparator)
	// 读取配置文件
	configPath := fmt.Sprintf("%s%sconfig%sserver.toml", rootDir, pathSeparator, pathSeparator)
	config, err := config.NewServerConfig(configPath)
	if err != nil || config == nil {
		logger.Fatalf("配置文件读取失败:%s", err.Error())
	}
	serverService.config = config
	// 初始化日志记录对象
	logPath := fmt.Sprintf("%s%slogs%sserver.log", rootDir, pathSeparator, pathSeparator)
	err = logs.New(logPath, config.Logger.Name, config.Logger.Verbose, config.Logger.SystemLog)
	if err != nil {
		logger.Errorf("初始化日志记录对象失败:%s\n", err.Error())
	}

	return &ServerService{}
}

// Start 开启运行服务
func (s *ServerService) Start() {
	logger.Errorf("测试错误日志")
}

// Stop 结束服务
func (s *ServerService) Stop() {
	err := logs.Close()
	if err != nil {
		log.Printf("配置文件关闭失败:%s\n", err.Error())
	}
}
