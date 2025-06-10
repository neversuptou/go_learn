package weather

import (
	"errors"
	"fmt"
	"io"
	"learnGO/geo"
	"net/http"
	"net/url"
	"strconv"
)

var ErrorWrongFormat = errors.New("WRONG_FORMAT")
var ErrorWrongUrl = errors.New("WRONG_URL")
var ErrorGetBody = errors.New("GET_BODY_ERRROR")
var ErrorStatusNot200 = errors.New("STATUS_CODE_NOT200")
var ErrorReadBody = errors.New("READ_BODY_ERRROR")

func GetWeather(geo geo.GeoData, format int) (string, error) {
	var result string
	if format < 1 || format > 4 {
		return "", ErrorWrongFormat
	}
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		fmt.Println(err.Error())
		return "", ErrorWrongUrl
	}
	params := url.Values{}
	params.Add("format", strconv.Itoa(format))
	baseUrl.RawQuery = params.Encode()

	response, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println(err)
		return "", ErrorGetBody
	}

	if response.StatusCode != 200 {
		return "NOT200", ErrorStatusNot200
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return "", ErrorReadBody
	}
	result = string(body)

	return result, nil
}
