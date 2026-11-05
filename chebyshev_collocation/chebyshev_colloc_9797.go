package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9797 evaluates chebyshev collocation node order 9797.
func ComputeChebyshevColloc9797(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
