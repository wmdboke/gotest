package weather

type Result struct {
	Err     error
	Weather string
	City    string
}

func NewResult(city, weather string, err error) *Result {
	ret := &Result{
		Err:     err,
		Weather: weather,
		City:    city,
	}
	return ret
}

// 获取城市
func (r *Result) GetCity() string {
	return r.City
}

// 获取天气
func (r *Result) GetWeather() string {
	return r.Weather
}

// 获取错误
func (r *Result) GetError() error {
	return r.Err
}
