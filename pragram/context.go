package pragram

import (
	"context"
	"log"
	"os"
	"time"
)

var (
	logg *log.Logger
)

func work(ctx context.Context, ch chan bool) {
	for {
		select {
		case <-ctx.Done():
			logg.Println(`下班!`)
			ch <- true
			return
		default:
			logg.Println(`上班!`)
			time.Sleep(2 * time.Second)
		}
	}
}

func ContextTest() {
	ch := make(chan bool)
	logg = log.New(os.Stdout, "", log.Ltime)
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	go work(ctx, ch)
	time.Sleep(2 * time.Second)
	//取消函数：当cancel被调用时,context.WithTimeout设置的时间超过后,关闭ctx.Done通道；
	logg.Println(`cancel!`)
	cancel()
	// 这个chan是为了保证子的goroutine执行完,当然也可以不用chan用time.Sleep停止几秒
	<-ch
	logg.Println(`无脑发呆中!`)
}
