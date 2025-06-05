package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"learnGO/geo"
)

func main() {
	city := flag.String("city", "", "City Name")
	//format := flag.Int("format", 1, "Format output the weather")
	flag.Parse()

	fmt.Println(*city)
	geo, err := geo.GetMyLocation(*city)
	if err != nil {
		color.Red(err.Error())
		return
	}
	fmt.Println(*geo)
}
