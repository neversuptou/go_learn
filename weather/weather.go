package weather

import (
	"fmt"
	"io"
	"learnGO/geo"
	"net/http"
	"net/url"
	"strconv"
)

func GetWeather(geo geo.GeoData, format int) string {
	result := ""
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	params := url.Values{}
	params.Add("format", strconv.Itoa(format))
	baseUrl.RawQuery = params.Encode()

	response, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println(err)
	}

	if response.StatusCode != 200 {
		return "NOT200"
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	result = string(body)

	return result
}
