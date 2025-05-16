package main

import (
	"ServiceB/infrastructure"
	"ServiceB/infrastructure/etcdx"
	"ServiceB/infrastructure/viperx"
	"github.com/gin-gonic/gin"
	"log"
)

type BServer struct {
	serve *gin.Engine
	conf  *infrastructure.AppConf
}

func NewBServer(conf *infrastructure.AppConf, serve *gin.Engine) *BServer {
	return &BServer{
		conf:  conf,
		serve: serve,
	}
}
func main() {
	// 1. 初始化 etcdx
	etcdx.InitEtcd()
	err := etcdx.PushConfigFromFileToEtcd("./infrastructure/config.yaml", etcdx.Cli)
	if err != nil {
		panic(err)

	}
	// 2. 初始化配置对象，并从 etcdx 拉配置
	setting := viperx.NewVipperSetting()

	// 从 etcdx 加载初始配置
	viperx.LoadConfigFromEtcd(setting)

	//3. 注入依赖
	s, err := InitService(setting)
	if err != nil {
		log.Fatalf("wire 初始化失败: %v", err)
	}

	s.Run()
}
func (s *BServer) Run() {
	s.serve.Run(s.conf.Addr)
}
