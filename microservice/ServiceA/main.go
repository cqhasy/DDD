package main

import (
	"ServiceA/application"
	"ServiceA/infrastructure"
	"ServiceA/infrastructure/etcdx"
	"ServiceA/infrastructure/viperx"
	"github.com/gin-gonic/gin"
	"log"
)

const BuyHandlerPath = "interface/ServiceA/BuyHandlerPath"
const GetInfoHandlerPath = "interface/ServiceA/GetInfoHandlerPath"

type Path struct {
	Path string `json:"path"`
}
type AServer struct {
	serve *application.Application
	conf  *infrastructure.AppConf
}

func NewAServer(conf *infrastructure.AppConf, serve *application.Application) *AServer {
	return &AServer{
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

	//	3. 注入依赖
	s, err := InitService(setting)
	if err != nil {
		log.Fatalf("wire 初始化失败: %v", err)
	}
	s.Run()
}
func (s *AServer) Run() {
	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/info/:id", s.serve.GetInfoHandler)
	r.GET("/buy/:id/:num", s.serve.BuyHandler)
	r.POST("/Add", s.serve.AddHandler)
	etcdx.RegisterInterface(BuyHandlerPath, Path{Path: s.conf.Addr + "/buy"})
	etcdx.RegisterInterface(GetInfoHandlerPath, Path{Path: s.conf.Addr + "/info"})
	r.Run(s.conf.Addr)

}
