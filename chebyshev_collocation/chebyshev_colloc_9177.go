package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9177 evaluates chebyshev collocation node order 9177.
func ComputeChebyshevColloc9177(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
