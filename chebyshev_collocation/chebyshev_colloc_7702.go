package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7702 evaluates chebyshev collocation node order 7702.
func ComputeChebyshevColloc7702(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
