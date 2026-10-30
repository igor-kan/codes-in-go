package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9232 evaluates chebyshev collocation node order 9232.
func ComputeChebyshevColloc9232(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
