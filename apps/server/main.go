package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/kardianos/service"
	"github.com/shiguanghuxian/athena/modules/common"
	"github.com/shiguanghuxian/athena/server"
)

// 程序版本
var (
	VERSION    string
	BUILD_TIME string
	GO_VERSION string
	GIT_HASH   string
)

var logger service.Logger
var serverService *server.ServerService

type program struct{}

func (p *program) Start(s service.Service) error {
	// 写时间戳到文件
	// 开启异步任务，开启服务
	go p.run()
	return nil
}

func (p *program) run() {
	// 存储pid
	err := common.WritePidToFile("server")
	if err != nil {
		log.Println("写pid文件错误")
	}
	// 实例化web服务
	serverService = server.New()
	serverService.Start()
}

func (p *program) Stop(s service.Service) error {
	// 删除pid文件
	common.RemovePidFile("server")
	// 关闭连接
	serverService.Stop()
	// 停止任务，3秒后返回
	<-time.After(time.Second * 1)
	return nil
}

func main() {
	// 全部核心运行程序
	runtime.GOMAXPROCS(runtime.NumCPU())

	svcConfig := &service.Config{
		Name:        "server.athena",
		DisplayName: "server.athena",
		Description: "Athena运维系统数据收集服务",
	}
	// 实例化
	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	// 接收一个参数，install | uninstall | start | stop | restart
	if len(os.Args) > 1 {
		if os.Args[1] == "-v" || os.Args[1] == "-version" {
			ver := fmt.Sprintf("Version: %s\nBuilt: %s\nGo version: %s\nGit commit: %s", VERSION, BUILD_TIME, GO_VERSION, GIT_HASH)
			fmt.Println(ver)
			return
		}
		err = service.Control(s, os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
