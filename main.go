package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type WeatherInfoJson struct {
	Weatherinfo WeatherinfoObject
}

type WeatherinfoObject struct {
	City    string
	CityId  string
	Temp    string
	WD      string
	WS      string
	SD      string
	WSE     string
	Time    string
	IsRadar string
	Radar   string
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	resp, err := http.Get("http://www.weather.com.cn/data/sk/101050101.html")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	input, err := ioutil.ReadAll(resp.Body)

	var jsonWeather WeatherInfoJson
	json.Unmarshal(input, &jsonWeather)
	w := jsonWeather.Weatherinfo
	fmt.Printf("%v现在%v度,有%v。\n", w.City, w.Temp, w.WD)
}
