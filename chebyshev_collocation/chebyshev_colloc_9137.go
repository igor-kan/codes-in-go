package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9137 evaluates chebyshev collocation node order 9137.
func ComputeChebyshevColloc9137(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
