package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9327 evaluates chebyshev collocation node order 9327.
func ComputeChebyshevColloc9327(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
