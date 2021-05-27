package weather

type WeatherHandler interface {
	// 并发个数
	SetCon(int) *FetchWeather
	// 设置回调
	SetHandler(func([]Result)) *FetchWeather
	// 设置超时时间
	SetTimeout(int) *FetchWeather
	// 设置城市
	SetCities([]string) *FetchWeather
	//获取天气
	GetWeather()
}

type ResultHandler interface {
	// 获取城市
	GetCity() string
	// 获取天气
	GetWeather() string
	// 获取错误
	GetError() error
}
