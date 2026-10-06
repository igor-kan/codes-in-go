package chebyshev_collocation

import "math"

// ComputeChebyshevColloc92 evaluates chebyshev collocation node order 92.
func ComputeChebyshevColloc92(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
