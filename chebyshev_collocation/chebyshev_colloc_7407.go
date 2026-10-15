package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7407 evaluates chebyshev collocation node order 7407.
func ComputeChebyshevColloc7407(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
