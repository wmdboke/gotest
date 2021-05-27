package etcdtest

import (
	"context"
	"fmt"
	etcdclient "github.com/coreos/etcd/client"
	"github.com/coreos/etcd/clientv3"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func EtcdClientV3(ctx context.Context) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect succ")
	defer cli.Close()
	//kv := clientv3.NewKV(clientv3)
	//ctx, cancleFunc := context.WithTimeout(context.TODO(), 5*time.Second)
	//
	//putResp, err := kv.Put(ctx, "/job/v3", "push the box", clientv3.WithPrevKV()) //withPrevKV()是为了获取操作前已经有的key-value
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%v", putResp.PrevKv)
	//
	//getResp, err := kv.Get(ctx, "/job/", clientv3.WithPrefix()) //withPrefix()是未了获取该key为前缀的所有key-value
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%v", getResp.Kvs)
	//
	//wc := client.Watch(context.Background(), "/job/", clientv3.WithPrefix(), clientv3.WithPrevKV())
	//for v := range wc {
	//	if v.Err() != nil {
	//		panic(err)
	//	}
	//	for _, e := range v.Events {
	//		fmt.Printf("type:%v\n kv:%v  prevKey:%v  ", e.Type, e.Kv, e.PrevKv)
	//	}
	//}
}

func EtcdClientConnect(ctx context.Context) {
	var DefaultEtcdTransport etcdclient.CancelableTransport = &http.Transport{
		//Dial:                dialer,
		//TLSClientConfig:     tlsConfig,
		TLSHandshakeTimeout: 10 * time.Second,
	}

	cfg := etcdclient.Config{
		Endpoints: []string{"localhost:2379"},
		Transport: DefaultEtcdTransport,
	}
	client, err := etcdclient.New(cfg)
	if err != nil {
		return
	}
	etcdclient.NewMembersAPI(client)
}


func CheckEtcd(endpoint string) error {
	ht := &http.Transport{}
	client := http.Client{
		Transport: ht,
	}
	defer ht.CloseIdleConnections()

	for i := 0; ; i++ {
		resp, err := client.Get(endpoint + "/health")
		if err != nil {
			if i > 1 {
				//logrus.Infof("Waiting on etcd startup: %v", err)
				fmt.Sprintf("Waiting on etcd startup: %v", err)
			}
			time.Sleep(time.Second)
			continue
		}
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			if i > 1 {
				//logrus.Infof("Waiting on etcd startup: status %d", resp.StatusCode)
				fmt.Sprintf("Waiting on etcd startup: status %d", resp.StatusCode)
			}
			time.Sleep(time.Second)
			continue
		}

		break
	}

	return nil
}
