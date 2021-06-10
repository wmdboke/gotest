package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go-test/gormtest"
	"go-test/pragram"
	"os"
	"strings"
	"time"

	//"path/filepath"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Action struct {
	Input  string `json:"input,omitempty"`
	Output string `json:"output,omitempty"`
}

// /Users/cwb/Documents/work/beeee/doc/rancher.log

func main() {
	fmt.Println(os.Args)

	//a := Action{Input: "aaa", Output: "bbb"}
	logrus.Infof("server Controller start sync_____%+v", os.Args)

	println("_____start my test progremmor： %s", "tttttt")

	name := "/v3/schemas/project/schemas/4444"
	parts := strings.SplitN(name, "/schemas/", 2)
	fmt.Println("strings.SplitN: part1 = %s, part2 = %s", name, parts[0])

	// select test
	selectX := pragram.Selectx{}
	selectX.Msg = make(chan int)
	ctx, canceler := context.WithCancel(context.Background())
	selectX.Ctx = ctx
	go selectX.SelectTest()
	time.Sleep(time.Second * 5)
	selectX.Msg <- 1
	canceler()
	select {}

	// websocket test

	// K8S 测试
	//k8stest.ConfTest()
	// 结构体测试
	//factory := grammartest.NewFactory()
	//factory.Name = "ssdf"

	// conf test
	//conftest.ReadConfig()
	//测试 Context
	// gostack.ContextTest()

	// 测试channel
	//gostack.ChannelTest()

	//os.Mkdir("management-state", 0700)

	// 测试 gorm
	gormtest.Gorm_test()

	// 测试错误吗
	//grammartest.Errors_test()

	// 运行命令行 以及 fork 进程
	//reexec_test()

	// 测试 etcd
	//etcdtest.CheckEtcd("http://192.168.137.245:2379")
	//etcdtest.EtcdClientConnect(context.Background())
}
