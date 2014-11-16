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

func rad(d float64) float64 {
	return d * math.Pi / 180
}

func Between(lat1 float64, lng1 float64, lat2 float64, lng2 float64) float64 { //refactor to object
	const meanRadius = 6371 //km
	lat1r, lat2r, lng1r, lng2r := rad(lat1), rad(lat2), rad(lng1), rad(lng2)
	x := (lng2r - lng1r) * math.Cos((lat1r+lat2r)/2)
	y := (lat2r - lat1r)
	return math.Sqrt(x*x+y*y) * meanRadius
}
