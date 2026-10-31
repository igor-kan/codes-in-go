package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9337 evaluates chebyshev collocation node order 9337.
func ComputeChebyshevColloc9337(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
