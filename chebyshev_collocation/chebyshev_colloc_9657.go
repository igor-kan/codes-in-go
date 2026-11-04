package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9657 evaluates chebyshev collocation node order 9657.
func ComputeChebyshevColloc9657(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
