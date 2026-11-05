package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9717 evaluates chebyshev collocation node order 9717.
func ComputeChebyshevColloc9717(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
