package main

import (
	"context"
	test "go-test/grpc/proto"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

func main() {
	// 建立连接到gRPC服务
	conn, err := grpc.Dial("127.0.0.1:8028", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 函数结束时关闭连接
	defer conn.Close()

	// 创建Waiter服务的客户端
	t := test.NewWaitClient(conn)

	// 模拟请求数据
	res := "test123"
	// os.Args[1] 为用户执行输入的参数 如：go run ***.go 123
	if len(os.Args) > 1 {
		res = os.Args[1]
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(1)*time.Second)

	// 调用gRPC接口
	tr, err := t.DoMD5(ctx, &test.Req{Fromjson: res})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("服务端响应: %s", tr.Md5Str)
}
