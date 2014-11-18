package distance

import "testing"


/* Shake Shack (40.74087, -73.98798) to Washington Sq (40.73146, -73.99709)
 * Should be about 1.3 km as the drone flies */

func TestBetween(t *testing.T) {
	a := location{lat: 40.74087, lng: -73.98798}
	b := location{lat: 40.73146, lng: -73.99709}
	const gt = 1.4
	const lt = 1.2
	if x := Between(a, b); x > gt || x < lt {
		t.Errorf("Between was %v, wanted something between %v and %v", x, lt, gt)
	}
}

/* Shake Shack (40.74087, -73.98798) to Harrah's (AC) (40.73146, -73.99709)
 * Should be about 155 km as the drone flies */

func TestBetweenLonger(t *testing.T) {
	a := location{lat: 40.74087, lng: -73.98798}
	b := location{lat: 39.38384, lng: -74.42712}
	const gt = 156
	const lt = 154
	if x := Between(a, b); x > gt || x < lt {
		t.Errorf("Between was %v, wanted something between %v and %v", x, lt, gt)
	}
}

//benchmarking the sample set

func BenchmarkBetween(bt *testing.B) {
	a := location{lat: 40.74087, lng: -73.98798}
	b := location{lat: 40.73146, lng: -73.99709}
	for n := 0; n < bt.N; n++ {
		Between(a, b)
	}
}
