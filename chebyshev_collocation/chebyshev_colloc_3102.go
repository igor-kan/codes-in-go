package chebyshev_collocation

import "math"

// ComputeChebyshevColloc3102 evaluates chebyshev collocation node order 3102.
func ComputeChebyshevColloc3102(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
