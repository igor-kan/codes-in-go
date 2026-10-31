package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9372 evaluates chebyshev collocation node order 9372.
func ComputeChebyshevColloc9372(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
