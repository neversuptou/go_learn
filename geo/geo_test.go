package geo_test

import (
	"errors"
	"learnGO/geo"
	"testing"
)

func TestGetMyLocation(t *testing.T) {
	// Arange - подготовка, excpected - результат, данные для функции
	сity := "London"
	expected := &geo.GeoData{
		City: "London",
	}

	// ACT -выполнение

	got, err := geo.GetMyLocation(сity)

	// ASSERT - проверка результата с expected

	if err != nil {
		t.Error("Error in get city")
	}
	if got.City != expected.City {
		t.Errorf("Expected %s, got %s", expected.City, got.City)
	}
}

func TestGetMyLocationNotCity(t *testing.T) {
	// Arange - подготовка, excpected - результат, данные для функции
	сity := "Moscow"

	// ACT -выполнение

	_, err := geo.GetMyLocation(сity)

	// ASSERT - проверка результата с expected

	if !errors.Is(err, geo.ErrorNoCity) {
		t.Errorf("Excepted %s, got %s", geo.ErrorNoCity, err)
	}
}
