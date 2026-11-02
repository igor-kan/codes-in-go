package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9457 evaluates chebyshev collocation node order 9457.
func ComputeChebyshevColloc9457(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
