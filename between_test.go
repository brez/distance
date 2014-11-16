package distance

import "testing"

// Shake Shack (40.74087, -73.98798) to Washington Sq (40.73146, -73.99709) 
// Should be 1.2976856489751614 km as the drone flies

func TestBetween(t *testing.T) {
  const lat1, lng1, lat2, lng2 = 40.74087, -73.98798, 40.73146, -73.99709
  const out = 1.2976856489751614
  if x := Between(lat1, lng1, lat2, lng2); x != out {
    t.Errorf("Between was %v, wanted %v", x, out)
  }
}