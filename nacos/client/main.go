package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/model"

	//"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/util"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func main() {
	sc := []constant.ServerConfig{
		{
			IpAddr: "127.0.0.1",
			Port:   8848,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         "e525eafa-f7d7-4029-83d9-008937f9d468", //namespace id
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}

	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}

	instanceParam := vo.RegisterInstanceParam{
		Ip:          "10.0.0.13",
		Port:        7777,
		ServiceName: "avServer.go",
		Weight:      10,
		ClusterName: "cluster",
		GroupName:   "group",
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
	}

	success, err := client.RegisterInstance(instanceParam)
	fmt.Printf("RegisterServiceInstance,param:%+v,result:%+v \n\n", instanceParam, success)
	if err != nil {
		fmt.Println("Register instance error: %s", err.Error())
	}

	scribeParam := &vo.SubscribeParam{
		ServiceName: "demo.go",
		Clusters:    []string{"cluster-b"},
		SubscribeCallback: func(services []model.SubscribeService, err error) {
			fmt.Printf("callback111 return services:%s \n\n", util.ToJsonString(services))
		},
	}
	client.Subscribe(scribeParam)

	select {}
}
