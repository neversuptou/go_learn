package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type GeoData struct {
	City   string `json:"city"`
	Format int    `json:"format"`
}

type cityResponse struct {
	Error bool `json:"error"`
}

func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		if !CheckCity(city) {
			return nil, errors.New("City not found")
		}
		return &GeoData{
			City: city,
		}, nil
	}
	response, err := http.Get("http://ipapi.co/json/")
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("NOT200")
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var geo GeoData
	err = json.Unmarshal(body, &geo)
	if err != nil {
		return nil, err
	}
	return &geo, nil
}

func CheckCity(city string) bool {
	body, _ := json.Marshal(map[string]string{
		"city": city,
	})

	resp, err := http.Post("https://countriesnow.space/api/v0.1/countries/population/cities", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	var geo cityResponse
	err = json.Unmarshal(body, &geo)
	return !geo.Error

}
