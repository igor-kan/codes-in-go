package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7917 evaluates chebyshev collocation node order 7917.
func ComputeChebyshevColloc7917(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
