package gostack

import "fmt"

func ChannelTest() {
	// 测试 channel
	println("start ChannelTest")
	ch := make(chan int, 4)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		// 如果不关闭channel,会引发panic
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
	println("end ChannelTest")
}
