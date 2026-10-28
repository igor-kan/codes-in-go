package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9172 evaluates chebyshev collocation node order 9172.
func ComputeChebyshevColloc9172(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
