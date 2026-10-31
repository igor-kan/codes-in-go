package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9357 evaluates chebyshev collocation node order 9357.
func ComputeChebyshevColloc9357(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
