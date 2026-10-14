package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7267 evaluates chebyshev collocation node order 7267.
func ComputeChebyshevColloc7267(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
