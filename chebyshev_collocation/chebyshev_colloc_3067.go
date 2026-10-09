package chebyshev_collocation

import "math"

// ComputeChebyshevColloc3067 evaluates chebyshev collocation node order 3067.
func ComputeChebyshevColloc3067(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
