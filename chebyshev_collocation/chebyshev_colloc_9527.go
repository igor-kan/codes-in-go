package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9527 evaluates chebyshev collocation node order 9527.
func ComputeChebyshevColloc9527(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
