package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9007 evaluates chebyshev collocation node order 9007.
func ComputeChebyshevColloc9007(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
