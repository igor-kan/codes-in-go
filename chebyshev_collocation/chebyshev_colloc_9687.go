package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9687 evaluates chebyshev collocation node order 9687.
func ComputeChebyshevColloc9687(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
