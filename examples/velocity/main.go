package main

import (
	"fmt"
	"time"

	"github.com/penguingovernor/length"
)

// Common velocities

// FeetPerSecond returns a velocity in meters per second.
func FeetPerSecond(l length.Distance, t time.Duration) float64 {
	ft := l / length.Feet
	s := t.Seconds()
	return float64(ft) / s
}

// MetersPerSecond returns a velocity in meters per second.
func MetersPerSecond(l length.Distance, t time.Duration) float64 {
	m := l / length.Meter
	s := t.Seconds()
	return float64(m) / s
}

// MilesPerHour returns a velocity in meters per second.
func MilesPerHour(l length.Distance, t time.Duration) float64 {
	mi := l / length.Mile
	h := t.Hours()
	return float64(mi) / h
}

// KilometersPerHour returns a velocity in meters per second.
func KilometersPerHour(l length.Distance, t time.Duration) float64 {
	km := l / length.Kilometer
	h := t.Hours()
	return float64(km) / h
}

func main() {

	// Let's say a slug on a rocket traveled 100.5 yards over the course of 3.2 seconds.
	distance, err := length.ParseDistance("100.5yd")
	if err != nil {
		panic(err)
	}
	t, err := time.ParseDuration("3.2s")
	if err != nil {
		panic(err)
	}

	type VelocityFunc func(length.Distance, time.Duration) float64
	funcUnitMap := map[string]VelocityFunc{
		"mi/h": MilesPerHour,
		"km/h": KilometersPerHour,
		"m/s":  MetersPerSecond,
		"ft/s": FeetPerSecond,
	}

	for k, v := range funcUnitMap {
		fmt.Println("Slug on rocket speed:", v(distance, t), k)
	}

	// Output:
	// 	Slug on rocket speed: 103.38435 km/h
	// Slug on rocket speed: 28.717875 m/s
	// Slug on rocket speed: 94.21875 ft/s
	// Slug on rocket speed: 64.24005681818181 mi/h
}
