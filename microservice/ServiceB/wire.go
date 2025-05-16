//go:generate wire
//go:build wireinject

package main

import (
	"ServiceB/application/service"
	"ServiceB/cmd/controller"
	"ServiceB/cmd/face"

	service1 "ServiceB/domain/service"
	"ServiceB/infrastructure"
	"ServiceB/infrastructure/cache"
	"ServiceB/infrastructure/dao"
	"ServiceB/infrastructure/viperx"
	"github.com/google/wire"
)

func InitService(setting *viperx.VipperSetting) (*BServer, error) {
	wire.Build(
		controller.ProvideApplication,
		service.ProvideService,
		service1.ProvideBRepository,
		infrastructure.NewCacheConf,
		infrastructure.NewAppConf,
		cache.InitCache,
		dao.NewDao,
		service1.NewService,
		service.NewApplication,
		controller.NewController,
		face.RegisterRoute,
		NewBServer,
	)
	return &BServer{}, nil
}
