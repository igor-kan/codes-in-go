package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9407 evaluates chebyshev collocation node order 9407.
func ComputeChebyshevColloc9407(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
