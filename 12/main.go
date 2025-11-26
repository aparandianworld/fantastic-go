package main

import (
	"fmt"
)

type Location struct {
	Latitude  float64
	Longitude float64
}

func NewLocation(lat float64, lon float64) (Location, error) {
	if lat < -90 || lat > 90 {
		return Location{}, fmt.Errorf("invalid latitude")
	}
	if lon < -180 || lon > 180 {
		return Location{}, fmt.Errorf("invalid longitude")
	}

	loc := Location{
		Latitude:  lat,
		Longitude: lon,
	}
	return loc, nil
}

func main() {
	latitude, longitude := 40.7128, -74.0060

	location := Location{Latitude: latitude, Longitude: longitude}

	fmt.Println(location)

	sf, err := NewLocation(37.7749, -122.4194)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(sf)
}
