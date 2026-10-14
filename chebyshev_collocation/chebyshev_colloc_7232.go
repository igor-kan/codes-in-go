package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7232 evaluates chebyshev collocation node order 7232.
func ComputeChebyshevColloc7232(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
