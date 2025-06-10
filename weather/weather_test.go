package weather_test

import (
	"errors"
	"learnGO/geo"
	"learnGO/weather"
	"strings"
	"testing"
)

var testCases = []struct {
	name   string
	format int
}{
	{name: "big format", format: -125},
	{name: "0 format", format: 0},
	{name: "minus format", format: -1},
}

func TestGetWeather(t *testing.T) {
	excepted := "London"
	var geodata = geo.GeoData{
		City:   excepted,
		Format: 4,
	}

	str, err := weather.GetWeather(geodata, geodata.Format)
	if err != nil {
		t.Errorf("Got %s", err)
	}
	if !strings.Contains(str, excepted) {
		t.Errorf("Excepted %s, got %s", excepted, str)
	}

}

func TestGetWeatherIncorrectFormat(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			excepted := "London"
			var geodata = geo.GeoData{
				City: excepted,
			}
			str, err := weather.GetWeather(geodata, tc.format)
			if !errors.Is(err, weather.ErrorWrongFormat) {
				t.Errorf("Excepted %s, got %s", weather.ErrorWrongFormat, str)
			}
		})
	}
}
