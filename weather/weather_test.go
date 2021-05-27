package weather

import (
	"fmt"
	"testing"
)

// 错误对应用层来说： 请求不到就等于超时了
// 底层错误，包括：服务器内部错误，请求错误，请求超时

// 应该加一个 自定义 err类型的
func GetWeatherByCities(results []Result, err error) {
	// 说明参数错误
	if err != nil {
		fmt.Println("参数错误: ", err.Error())
		return
	}
	// 说明服务器没拿到值
	if len(results) == 0 {
		fmt.Println("UI 友好提示: timeout")
		return
	}

	for _, ret := range results {
		fmt.Println("[Result]:", ret.GetCity()+"天气:", ret.GetWeather())
	}
	fmt.Println("[Result]: 获取到天气情况的城市个数： ", len(results))
}

func TestWeather_GetWeather(t *testing.T) {
	// 查询天气
	cities := []string{"1-北京", "2-上海", "3-广州", "4-深圳", "5-天津", "6-杭州", "7-济南", "8-青岛", "9-苏州", "10-南京", "11-成都", "12-重庆"}
	w := NewFetchWeather()
	w.SetCon(4).SetTimeout(5).SetCities(cities).SetHandler(GetWeatherByCities)
	//w.SetCon(4).SetTimeout(10).SetHandler(GetWeatherByCities)

	w.GetWeather()

	//select {} // 验证协程退出锁

}
