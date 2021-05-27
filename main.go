package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go-test/gormtest"
	"os"
	"strings"

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
	logrus.Infof("server Controller start sync_____%+v", nil)

	println("_____start my test progremmor： %s", "tttttt")

	name := "/v3/schemas/project/schemas/4444"
	parts := strings.SplitN(name, "/schemas/", 2)
	name = parts[1]

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

	//力扣
	//leetcode.LongestPalindrome("adfbg")

	// 运行命令行 以及 fork 进程
	//reexec_test()

	// 测试 etcd
	//etcdtest.CheckEtcd("http://192.168.137.245:2379")
	//etcdtest.EtcdClientConnect(context.Background())
}
