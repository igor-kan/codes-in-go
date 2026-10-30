package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9267 evaluates chebyshev collocation node order 9267.
func ComputeChebyshevColloc9267(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
