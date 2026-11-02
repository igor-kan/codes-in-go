package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9497 evaluates chebyshev collocation node order 9497.
func ComputeChebyshevColloc9497(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
