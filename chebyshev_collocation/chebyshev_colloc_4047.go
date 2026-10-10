package chebyshev_collocation

import "math"

// ComputeChebyshevColloc4047 evaluates chebyshev collocation node order 4047.
func ComputeChebyshevColloc4047(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
