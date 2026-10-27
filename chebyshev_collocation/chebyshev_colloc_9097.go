package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9097 evaluates chebyshev collocation node order 9097.
func ComputeChebyshevColloc9097(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
