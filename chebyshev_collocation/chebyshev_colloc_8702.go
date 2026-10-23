package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8702 evaluates chebyshev collocation node order 8702.
func ComputeChebyshevColloc8702(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
