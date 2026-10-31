package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9307 evaluates chebyshev collocation node order 9307.
func ComputeChebyshevColloc9307(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
