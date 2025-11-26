package main

import (
	"fmt"
)

type Location struct {
	Latitude  float64
	Longitude float64
}

// Embedding Location in Car
type Car struct {
	id string
	Location
}

func NewCar(id string, lat float64, lon float64) (Car, error) {

	loc, err := NewLocation(lat, lon)
	if err != nil {
		return Car{}, err
	}

	car := Car{
		id:       id,
		Location: loc,
	}

	return car, nil
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

func (l *Location) moveLocation(lat float64, lon float64) *Location {
	l.Latitude += lat
	l.Longitude += lon

	return l
}

func main() {
	// TODO: implement main

	car, err := NewCar("Tesla Model X", 40.7128, -74.0060)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
		return
	}

	car.moveLocation(20.07, -10.9)

	fmt.Printf("%+v\n", car)
}
