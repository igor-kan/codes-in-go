package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8917 evaluates chebyshev collocation node order 8917.
func ComputeChebyshevColloc8917(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
