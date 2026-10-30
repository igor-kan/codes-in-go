package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9222 evaluates chebyshev collocation node order 9222.
func ComputeChebyshevColloc9222(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
