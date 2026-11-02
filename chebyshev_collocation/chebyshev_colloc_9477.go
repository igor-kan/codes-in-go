package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9477 evaluates chebyshev collocation node order 9477.
func ComputeChebyshevColloc9477(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
