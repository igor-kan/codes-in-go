package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9257 evaluates chebyshev collocation node order 9257.
func ComputeChebyshevColloc9257(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
