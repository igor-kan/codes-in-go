package chebyshev_collocation

import "math"

// ComputeChebyshevColloc62 evaluates chebyshev collocation node order 62.
func ComputeChebyshevColloc62(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
