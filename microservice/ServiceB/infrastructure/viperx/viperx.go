package viperx

import (
	"bytes"
	"context"

	"ServiceB/infrastructure/etcdx"
	"github.com/spf13/viper"
	"log"
)

const EtcdConfigKey = "/config/serviceB/config.yaml"

type VipperSetting struct {
	vp *viper.Viper
}

func NewVipperSetting() *VipperSetting {
	v := viper.New()
	v.SetConfigType("yaml")
	return &VipperSetting{vp: v}
}

func (s *VipperSetting) ReadSection(k string, v interface{}) error {
	return s.vp.UnmarshalKey(k, v)
}

func (s *VipperSetting) LoadFromBytes(data []byte) error {
	return s.vp.ReadConfig(bytes.NewBuffer(data))
}

func (s *VipperSetting) WatchEtcdAndReload() {
	go func() {
		rch := etcdx.Cli.Watch(context.Background(), EtcdConfigKey)
		for wresp := range rch {
			for _, ev := range wresp.Events {
				log.Printf("配置变更: %s %q : %q", ev.Type, ev.Kv.Key, ev.Kv.Value)
				err := s.LoadFromBytes(ev.Kv.Value)
				if err != nil {
					log.Printf("热加载失败: %v", err)
				} else {
					log.Println("配置热加载完成")
				}
			}
		}
	}()
}

func LoadConfigFromEtcd(setting *VipperSetting) {
	resp, err := etcdx.Cli.Get(context.Background(), EtcdConfigKey)
	if err != nil || len(resp.Kvs) == 0 {
		log.Fatalf("无法从 etcdx 获取配置: %v", err)
	}
	err = setting.LoadFromBytes(resp.Kvs[0].Value)
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}
	log.Println("初始化配置完成:", setting.vp.AllSettings())

	setting.WatchEtcdAndReload()
}
