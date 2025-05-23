// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"ServiceA/application"
	"ServiceA/domain/service"
	"ServiceA/infrastructure"
	"ServiceA/infrastructure/cache"
	"ServiceA/infrastructure/dao"
	"ServiceA/infrastructure/viperx"
)

// Injectors from wire.go:

func InitService(setting *viperx.VipperSetting) (*AServer, error) {
	appConf := infrastructure.NewAppConf(setting)
	cacheConfig := infrastructure.NewCacheConf(setting)
	client := cache.InitCache(cacheConfig)
	daoDao := dao.NewDao(client)
	aService := service.NewService(daoDao)
	applicationApplication := application.NewApplication(aService)
	aServer := NewAServer(appConf, applicationApplication)
	return aServer, nil
}
