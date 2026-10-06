package chebyshev_collocation

import "math"

// ComputeChebyshevColloc102 evaluates chebyshev collocation node order 102.
func ComputeChebyshevColloc102(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
