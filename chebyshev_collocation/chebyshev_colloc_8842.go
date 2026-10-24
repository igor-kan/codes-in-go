package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8842 evaluates chebyshev collocation node order 8842.
func ComputeChebyshevColloc8842(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
