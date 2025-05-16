//go:generate wire
//go:build wireinject

package main

import (
	"ServiceA/application"
	"ServiceA/domain/service"
	"ServiceA/infrastructure"
	"ServiceA/infrastructure/cache"
	"ServiceA/infrastructure/dao"
	"ServiceA/infrastructure/viperx"
	"github.com/google/wire"
)

func InitService(setting *viperx.VipperSetting) (*AServer, error) {
	wire.Build(
		infrastructure.NewCacheConf,
		infrastructure.NewAppConf,
		cache.InitCache,
		dao.NewDao,
		service.NewService,
		application.NewApplication,
		NewAServer,
	)
	return &AServer{}, nil
}
