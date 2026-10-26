package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9027 evaluates chebyshev collocation node order 9027.
func ComputeChebyshevColloc9027(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
