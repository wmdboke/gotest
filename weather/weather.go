package weather

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"sync"
	"time"
)

const delay = 3

type CallBackFunc func([]Result, error)

type FetchWeather struct {
	cities   []string
	callback CallBackFunc
	timeout  time.Duration
	con      int

	// 控制退出的
	isClosed bool
	sync.RWMutex
}

// 给默认值
func NewFetchWeather() *FetchWeather {
	ret := &FetchWeather{
		timeout:  4,  // 默认值走配置文件
		con:      10, // 默认值走配置文件
		isClosed: false,
	}
	return ret
}

// 并发个数
func (w *FetchWeather) SetCon(con int) *FetchWeather {
	w.con = con
	return w
}

// 设置回调
func (w *FetchWeather) SetHandler(call CallBackFunc) *FetchWeather {
	w.callback = call
	return w
}

// 设置超时时间
func (w *FetchWeather) SetTimeout(timeout time.Duration) *FetchWeather {
	w.timeout = timeout
	return w
}

// 设置城市
func (w *FetchWeather) SetCities(cities []string) *FetchWeather {
	w.cities = cities
	return w
}

// 单个城市请求，，http请求加入超时控制
func (w *FetchWeather) fetchOne(ctx context.Context, baseUrl string, city string, timeout time.Duration) (*Result, error) {
	params := url.Values{}
	params.Set("city", city)
	req, err := NewRequest("GET", baseUrl, params)
	if err != nil {
		return nil, err
	}
	_, err = Do(ctx, req, timeout)
	time.Sleep(time.Second * delay) //模拟超时
	// 屏蔽掉验证正常的情况
	//if err != nil {
	//	return nil, err
	//}

	// 解析byte
	// parse

	// 解析完数据组装Result
	ret := NewResult(city, "晴朗", nil)
	return ret, nil
}

// 获取天气的执行模块, 循环请求
func (w *FetchWeather) fetch(ctx context.Context, cities []string, timeout time.Duration, c chan *Result, idx int) {
	citiesNum := len(cities)
	if citiesNum <= 0 {
		return
	}

	for _, city := range cities {
		// 循环请求城市，每个给一定的时间去获取天气,如果有一个报错了，就不继续请求了
		one, err := w.fetchOne(ctx, "baseUrl", city, timeout/time.Duration(citiesNum))
		if err != nil {
			fmt.Println("[Service]: fetch city error", err)
			continue
		}
		fmt.Println("[Service]: ", idx, " fetch city ", city, " ok")

		// 如果关闭了那么直接退出，不需要继续调用了
		w.RLock()
		if w.isClosed == true {
			fmt.Println("[Service]: ", idx, " timeout ")
			w.RUnlock()
			return
		}
		w.RUnlock()

		// 退出不需要发送了
		c <- one
	}

	return
}

// 获取天气
func (w *FetchWeather) GetWeather() {
	var (
		results []Result
	)

	citiesNum := len(w.cities)
	if citiesNum <= 0 {
		fmt.Println("没有设置要查询的城市列表")
		err := errors.New("没有设置要查询的城市列表")
		w.callback(results, err)
		return
	}
	startTime := time.Now()
	defer func() {
		w.Lock()
		defer w.Unlock()
		w.isClosed = true

		w.callback(results, nil)

		endTime := time.Now()
		costTime := endTime.Sub(startTime).Seconds()
		fmt.Printf("fetch weather over. cost time = %fs \n", costTime)
	}()

	ctx := context.Background()
	timeout := w.timeout * time.Second
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	//num := (len(w.cities) + w.con - 1) / w.con
	num := citiesNum / w.con
	last := citiesNum - (num * w.con)
	chanResult := make(chan *Result, citiesNum)
	defer close(chanResult)

	fmt.Println("loop = ", w.con, "timeout = ", timeout, "num = ", num, "last = ", last)
	// 协程：启动队列查询
	for p := 0; p < last; p++ {
		arr := w.cities[p*(num+1) : (p+1)*(num+1)]
		go w.fetch(ctx, arr, w.timeout, chanResult, p)
	}

	for l := last; l < w.con; l++ {
		arr := w.cities[l*num : (l+1)*num]
		go w.fetch(ctx, arr, w.timeout, chanResult, l)
	}

	for {
		select {
		case ret := <-chanResult:
			results = append(results, *ret)
			if len(results) == len(w.cities) {
				return
			}

		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		}
	}
}
