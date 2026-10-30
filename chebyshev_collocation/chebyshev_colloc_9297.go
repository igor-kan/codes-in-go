package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9297 evaluates chebyshev collocation node order 9297.
func ComputeChebyshevColloc9297(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
