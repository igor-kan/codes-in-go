package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9607 evaluates chebyshev collocation node order 9607.
func ComputeChebyshevColloc9607(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
