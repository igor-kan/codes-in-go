package chebyshev_collocation

import "math"

// ComputeChebyshevColloc3047 evaluates chebyshev collocation node order 3047.
func ComputeChebyshevColloc3047(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
