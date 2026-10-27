package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9087 evaluates chebyshev collocation node order 9087.
func ComputeChebyshevColloc9087(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
