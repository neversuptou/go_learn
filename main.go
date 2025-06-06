package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"learnGO/geo"
	"learnGO/weather"
)

func main() {
	city := flag.String("city", "", "City Name")
	format := flag.Int("format", 2, "Format output the weather")
	flag.Parse()
	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		color.Red(err.Error())
		return
	}

	getWeather := weather.GetWeather(*geoData, *format)
	fmt.Println(getWeather)
}
