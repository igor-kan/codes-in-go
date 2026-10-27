package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9067 evaluates chebyshev collocation node order 9067.
func ComputeChebyshevColloc9067(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
