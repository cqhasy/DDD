package etcdx

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/spf13/viper"
	"go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

// todo:config文件发生变化时推送到etcd.
// todo:可以用[]byte来扩大使用范围。
type PathRemote struct {
	Path string `json:"path"`
}

const EtcdConfigKey = "/config/serviceB/config.yaml"

var Cli *clientv3.Client

func InitEtcd() {
	var err error
	Cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"60.205.12.92:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf("连接 etcd 失败: %v", err)
	}
}
func PushConfigFromFileToEtcd(configFilePath string, cli *clientv3.Client) error {
	// 使用viper读取配置文件
	// 使用 viper 读取配置文件
	viper.SetConfigFile(configFilePath)
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("无法读取配置文件: %w", err)
	}

	// 将 viper 配置转换为 YAML 字符串
	configData := viper.AllSettings()

	// 将配置内容转换为 YAML 格式
	configBytes, err := yaml.Marshal(configData)
	if err != nil {
		return fmt.Errorf("无法将配置内容转换为YAML格式: %w", err)
	}

	// 检查 etcd 中是否已经有配置
	resp, err := cli.Get(context.Background(), EtcdConfigKey)
	if err != nil {
		log.Fatalf("无法从 etcd 获取配置: %v", err)
		return err
	}

	// 如果配置不存在，将配置内容推送到 etcd
	if len(resp.Kvs) == 0 {
		// 将配置内容推送到 etcd
		_, err := cli.Put(context.Background(), EtcdConfigKey, string(configBytes))
		if err != nil {
			log.Fatalf("推送配置到 etcd 失败: %v", err)
			return err
		}
		log.Println("配置已推送到 etcd")
	} else {
		log.Println("配置已存在，无需推送配置")
	}

	return nil
}
func FindService(path string, config PathRemote) (PathRemote, error) {
	resp, err := Cli.Get(context.Background(), path)
	if err != nil {
		log.Fatal("Failed to get key:", err)
		return PathRemote{}, err
	}

	if len(resp.Kvs) > 0 {
		err = json.Unmarshal(resp.Kvs[0].Value, &config)

		if err != nil {
			log.Fatal("Failed to unmarshal value:", err)
			return PathRemote{}, err
		}

		return config, nil
	}
	return PathRemote{}, nil
}
func RegisterInterface(path string, in any) error {
	value, err := json.Marshal(in)
	if err != nil {
		return fmt.Errorf("failed to marshal value: %w", err)
	}
	_, err = Cli.Put(context.Background(), path, string(value))
	if err != nil {
		return err
	}
	return nil
}
