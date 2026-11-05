package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9792 evaluates chebyshev collocation node order 9792.
func ComputeChebyshevColloc9792(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
