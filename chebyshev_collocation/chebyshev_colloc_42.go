package chebyshev_collocation

import "math"

// ComputeChebyshevColloc42 evaluates chebyshev collocation node order 42.
func ComputeChebyshevColloc42(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
