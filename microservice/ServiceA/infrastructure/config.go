package infrastructure

import (
	"ServiceA/infrastructure/viperx"
)

type CacheConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}
type AppConf struct {
	Addr string `yaml:"addr"`
}

func NewCacheConf(s *viperx.VipperSetting) *CacheConfig {
	var cacheConf = &CacheConfig{}
	err := s.ReadSection("cache", cacheConf)
	if err != nil {
		return nil
	}
	return cacheConf
}
func NewAppConf(s *viperx.VipperSetting) *AppConf {
	var appConf = &AppConf{}
	err := s.ReadSection("app", appConf)
	if err != nil {
		return nil
	}
	return appConf
}
