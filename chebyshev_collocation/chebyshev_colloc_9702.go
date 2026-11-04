package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9702 evaluates chebyshev collocation node order 9702.
func ComputeChebyshevColloc9702(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
