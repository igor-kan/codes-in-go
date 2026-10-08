package chebyshev_collocation

import "math"

// ComputeChebyshevColloc1047 evaluates chebyshev collocation node order 1047.
func ComputeChebyshevColloc1047(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
