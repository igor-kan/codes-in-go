package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9437 evaluates chebyshev collocation node order 9437.
func ComputeChebyshevColloc9437(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
