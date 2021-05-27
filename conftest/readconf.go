package conftest

import (
	gcfg "gopkg.in/gcfg.v1"
	"os"
)

var filepath = "/Users/cwb/Documents/work/beeee/go/src/go-test/conftest/test.conf"

type Global struct { //配置文件要通过tag来指定配置文件中的名称
	MqttHostname  string `gcfg:"mqtt-hostname"`
	MqttPort      string `gcfg:"mqtt-port"`
	MqttUser      string `gcfg:"mqtt-user"`
	MqttPass      string `gcfg:"mqtt-pass"`
	MqttKeepalive int    `gcfg:"mqtt-keepalive" `
	MqttTimeout   int    `gcfg:"mqtt-timeout"`
}

type Config struct { //配置文件要通过tag来指定配置文件中的名称
	Global
}

////读取配置文件并转成结构体
//func ReadConfig() (Config, error) {
//	ReadConfigCinderProvider()
//
//	var config Config
//	conf, err := ini.Load(filepath) //加载配置文件
//	if err != nil {
//		log.Println("load config file fail!")
//		return config, err
//	}
//	conf.BlockMode = false
//	err = conf.MapTo(&config) //解析成结构体
//	if err != nil {
//		log.Println("mapto config file fail!")
//		return config, err
//	}
//	log.Println(config)
//	return config, nil
//}

func ReadConfig() error {
	var cfg Config
	config, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer config.Close()

	err = gcfg.FatalOnly(gcfg.ReadInto(&cfg, config))
	if err != nil {
		return err
	}

	return nil
}
