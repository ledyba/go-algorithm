package algo

import (
	"math"
	"testing"
)

func compareSlice(a, b []int) bool {
	if len(a) == 0 && len(b) == 0 {
		return true
	} else if len(a) != len(b) {
		return false
	}
	return a[0] == b[0] && compareSlice(a[1:], b[1:])
}

func TestBitonic(t *testing.T) {
	pts := []Point{Point{0, 6}, Point{1, 0}, Point{2, 3}, Point{5, 4}, Point{6, 1}, Point{7, 5}, Point{8, 2}}
	route, score := BitonicTour(pts)
	if math.Abs(float64(score)-25.58) > 0.1 {
		t.Fatalf("Too Far %f", score)
	}
	ans := []int{5, 3, 2, 0, 1, 4, 6}
	if !compareSlice(route, ans) {
		t.Fatalf("%v != %v", route, ans)
	}
}
