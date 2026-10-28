package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9157 evaluates chebyshev collocation node order 9157.
func ComputeChebyshevColloc9157(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
