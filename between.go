package distance

import "math"

/*
 * Distance Between Two Points
 *
 * Equirectangular approximation (Pythagoras' theorem) because we'd rather have speed than accuracy in short distances.
 *
 * x = Δλ ⋅ cos φm
 * y = Δφ
 * d = R ⋅ √(x² + y²)
 *
 * Where φ is latitude, λ is longitude, and R is the earth's mean radius
 *
 */

const meanRadius = 6371 //km

type location struct {
	lat float64
	lng float64
}

func rad(d float64) float64 {
	return d * math.Pi / 180
}

func Between(a location, b location) float64 {
	latXr, latYr, lngXr, lngYr := rad(a.lat), rad(b.lat), rad(a.lng), rad(b.lng)
	x := (lngYr - lngXr) * math.Cos((latXr+latYr)/2)
	y := (latYr - latXr)
	return math.Sqrt(x*x+y*y) * meanRadius
}
